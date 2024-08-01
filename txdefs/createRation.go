package txdefs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/events"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

// CreateRation transaction definition
var CreateRation = tx.Transaction{
	Tag:         "createRation",
	Label:       "Create Ration",
	Description: "Create a new ration",
	Method:      "POST",
	Callers: []accesscontrol.Caller{ // Only org2 admin can call this transaction
		{
			MSP: "org2MSP",
			OU:  "admin",
		},
		{
			MSP: "orgMSP",
			OU:  "admin",
		},
	},
	// Define the arguments required for this transaction
	Args: []tx.Argument{
		{
			Tag:         "id",
			Label:       "Ration ID",
			Description: "Ration ID",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "category",
			Label:       "Ration Category",
			Description: "Ration Category",
			DataType:    "rationCategory",
			Required:    true,
		},
		{
			Tag:         "description",
			Label:       "Ration Description",
			Description: "Ration Description",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "package",
			Label:       "Ration Package",
			Description: "Ration Package",
			DataType:    "packageType",
			Required:    true,
		},
		{
			Tag:         "distributedBy",
			Label:       "Distributed By",
			Description: "Distributed By",
			DataType:    "->distributionPoint",
			Required:    false,
		},
		{
			Tag:         "quantity",
			Label:       "Ration Quantity",
			Description: "Ration Quantity",
			DataType:    "integer",
			Required:    true,
		},
		{
			Tag:         "expiryDate",
			Label:       "Expiry Date",
			Description: "Expiry Date",
			DataType:    "datetime",
			Required:    true,
		},
		{
			Tag:         "mfgDate",
			Label:       "Manufacturing Date",
			Description: "Manufacturing Date",
			DataType:    "datetime",
			Required:    true,
		},
		{
			Tag:         "batchNumber",
			Label:       "Batch Number",
			Description: "Batch Number",
			DataType:    "integer",
			Required:    true,
		},
	},
	// Define the method to be executed for this transaction
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		// Get the values from the request
		id, _ := req["id"].(string)
		category, _ := req["category"].(string)
		description, _ := req["description"].(string)
		rationPackage, _ := req["package"].(string)
		distributedBy, _ := req["distributedBy"].(assets.Key)
		quantity, _ := req["quantity"].(int)
		expiryDate, _ := req["expiryDate"].(time.Time)
		mfgDate, _ := req["mfgDate"].(time.Time)
		batchNumber, _ := req["batchNumber"].(int)

		// Check if the ration package is valid
		rationMap := make(map[string]interface{})
		rationMap["@assetType"] = "ration"
		rationMap["id"] = id
		rationMap["category"] = category
		rationMap["description"] = description
		rationMap["package"] = rationPackage

		rationMap["quantity"] = quantity
		rationMap["expiryDate"] = expiryDate
		rationMap["mfgDate"] = mfgDate
		rationMap["batchNumber"] = batchNumber

		// get distributedBy asset
		if distributedBy != nil {
			d, err := distributedBy.Get(stub)
			if err != nil {
				return nil, errors.WrapError(err, "failed to get distributedBy asset from the ledger")
			}
			rationMap["distributedBy"] = d
		}

		// Create a new ration asset
		rationAsset, err := assets.NewAsset(rationMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new ration asset on the blockchain
		_, err = rationAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}
		// Marshal the asset back to JSON format
		rationJSON, nerr := json.Marshal(rationAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		// Log the event
		logMsg, ok := json.Marshal(fmt.Sprintf("New ration created: %s", id))
		if ok != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		events.CallEvent(stub, "rationCreatedLog", logMsg)

		return rationJSON, nil
	},
}
