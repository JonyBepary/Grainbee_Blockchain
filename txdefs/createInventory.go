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

var CreateInventory = tx.Transaction{
	Tag:         "createInventory",
	Label:       "Create Inventory",
	Description: "Create a new inventory",
	Method:      "POST",
	Callers: []accesscontrol.Caller{ // Only org3 admin can call this transaction
		{
			MSP: "org3MSP",
			OU:  "admin",
		},
		{
			MSP: "orgMSP",
			OU:  "admin",
		},
	},
	Args: []tx.Argument{
		{
			Tag:         "name",
			Label:       "Inventory Name",
			Description: "Inventory Name",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "rations",
			Label:       "Ration Collection",
			Description: "Ration Collection",
			DataType:    "[]->ration",
			Required:    false,
		},
		{
			Tag:         "entranceCode",
			Label:       "Entrance Code for the Inventory",
			Description: "Entrance Code for the Inventory",
			DataType:    "->secret",
			Required:    false,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["name"].(string)
		rations, _ := req["rations"].([]interface{})
		entranceCode, _ := req["entranceCode"].(string)

		inventoryMap := make(map[string]interface{})
		inventoryMap["@assetType"] = "inventory"
		inventoryMap["name"] = name
		inventoryMap["rations"] = rations
		inventoryMap["entranceCode"] = entranceCode

		inventoryAsset, err := assets.NewAsset(inventoryMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = inventoryAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		inventoryJSON, nerr := json.Marshal(inventoryAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		logMsg, ok := json.Marshal(fmt.Sprintf("New inventory created: %s", name))
		if ok != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		events.CallEvent(stub, "inventoryCreatedLog", logMsg)

		return inventoryJSON, nil
	},
}
