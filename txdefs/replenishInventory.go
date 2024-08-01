package txdefs

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/events"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var ReplenishInventory = tx.Transaction{
	Tag:         "replenishInventory",
	Label:       "Replenish Inventory",
	Description: "Replenish the inventory of a distribution point",
	Method:      "POST",
	Callers: []accesscontrol.Caller{ // Only org1 admin can call this transaction
		{
			MSP: "org1MSP",
			OU:  "admin",
		},
		{
			MSP: "orgMSP",
			OU:  "admin",
		},
	},
	Args: []tx.Argument{
		{
			Tag:         "distributionPointId",
			Label:       "Distribution Point ID",
			Description: "Distribution Point ID",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "rations",
			Label:       "Rations",
			Description: "List of rations to add to the inventory",
			DataType:    "[]->ration",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		distributionPointId, _ := req["distributionPointId"].(string)
		rations, _ := req["rations"].([]interface{})

		// Retrieve the distribution point asset
		distributionPointKey, ok := req["distributionPointId"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "failed to get distribution point ID from the request")
		}

		distributionPointAsset, err := distributionPointKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get distribution point asset from the ledger")
		}

		// Update the distribution point asset with the new rations
		distributionPointMap := (map[string]interface{})(*distributionPointAsset)
		if distributionPointMap["inventory"] == nil {
			distributionPointMap["inventory"] = make(map[string]interface{})
		}
		inventoryMap := distributionPointMap["inventory"].(map[string]interface{})

		for _, ration := range rations {
			rationMap := ration.(map[string]interface{})
			rationId := rationMap["id"].(string)
			inventoryMap[rationId] = ration
		}

		// Concurrently update ration assets
		var wg sync.WaitGroup
		var mutex sync.Mutex
		var updateErr error

		for _, ration := range rations {
			wg.Add(1)
			go func(ration interface{}) {
				defer wg.Done()

				rationKey, ok := ration.(assets.Key)
				if !ok {
					mutex.Lock()
					updateErr = errors.WrapError(nil, "failed to get ration ID from the request")
					mutex.Unlock()
					return
				}

				rationAsset, err := rationKey.Get(stub)
				if err != nil {
					mutex.Lock()
					updateErr = errors.WrapError(err, "failed to get ration asset from the ledger")
					mutex.Unlock()
					return
				}

				rationMap := (map[string]interface{})(*rationAsset)
				rationMap["quantity"] = ration.(map[string]interface{})["quantity"]
				rationMap["expiryDate"] = ration.(map[string]interface{})["expiryDate"]
				rationMap["mfgDate"] = ration.(map[string]interface{})["mfgDate"]
				rationMap["batchNumber"] = ration.(map[string]interface{})["batchNumber"]
				rationMap["description"] = ration.(map[string]interface{})["description"]
				rationMap["package"] = ration.(map[string]interface{})["package"]
				rationMap["distributedBy"] = ration.(map[string]interface{})["distributedBy"]

				_, err = rationAsset.Update(stub, rationMap)
				if err != nil {
					mutex.Lock()
					updateErr = errors.WrapError(err, "failed to update ration asset")
					mutex.Unlock()
					return
				}
			}(ration)
		}

		wg.Wait()

		if updateErr != nil {
			return nil, errors.WrapError(updateErr, "failed to update ration asset")
		}
		distributionPointMap["inventory"] = inventoryMap

		updatedDistributionPointAsset, err := distributionPointAsset.Update(stub, distributionPointMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update distribution point asset")
		}

		// Marshal asset back to JSON format
		updatedDistributionPointJSON, nerr := json.Marshal(updatedDistributionPointAsset)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		// Marshal message to be logged
		logMsg, erre := json.Marshal(fmt.Sprintf("Inventory replenished for distribution point with ID: %s", distributionPointId))
		if erre != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		events.CallEvent(stub, "inventoryReplenishedLog", logMsg)

		return updatedDistributionPointJSON, nil
	},
}
