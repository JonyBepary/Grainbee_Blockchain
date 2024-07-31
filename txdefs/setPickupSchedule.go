package txdefs

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/events"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var SetPickupSchedule = tx.Transaction{
	Tag:         "setPickupSchedule",
	Label:       "Set Pickup Schedule",
	Description: "Set the pickup schedule for a specific distribution point",
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
			Tag:         "pickupSchedule",
			Label:       "Pickup Schedule",
			Description: "Pickup Schedule",
			DataType:    "rationPickupSchedule",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		distributionPointId, _ := req["distributionPointId"].(string)
		pickupSchedule, _ := req["pickupSchedule"].(string)

		// Retrieve the distribution point asset
		distributionPointKey, ok := req["distributionPointId"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "failed to get distribution point ID from the request")
		}

		distributionPointAsset, err := distributionPointKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get distribution point asset from the ledger")
		}

		// Check if the schedule is colliding with any other schedule
		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"@assetType":          "distributionPoint",
				"pickupSchedule":      pickupSchedule,
				"distributionPointId": distributionPointId,
			},
		}

		response, err := assets.Search(stub, query, "", true)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error searching for distribution point's pickup schedule", 500)
		}

		if response.Metadata.XXX_Size() > 0 {
			return nil, errors.NewCCError("pickup schedule is colliding with another schedule", 400)
		}

		// Update the distribution point asset with the pickup schedule
		distributionPointMap := (map[string]interface{})(*distributionPointAsset)
		distributionPointMap["pickupSchedule"] = pickupSchedule

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
		logMsg, erre := json.Marshal(fmt.Sprintf("Pickup schedule set for distribution point: %s", distributionPointId))
		if erre != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		events.CallEvent(stub, "pickupScheduleSetLog", logMsg)

		return updatedDistributionPointJSON, nil
	},
}
