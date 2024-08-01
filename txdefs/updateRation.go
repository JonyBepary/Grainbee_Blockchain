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

var UpdateRation = tx.Transaction{
	Tag:         "updateRation",
	Label:       "Update Ration",
	Description: "Update the information of a ration",
	Method:      "PUT",
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
			Required:    false,
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
			DataType:    "string",
			Required:    true,
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
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, _ := req["id"].(string)

		// Retrieve the ration asset
		rationKey, ok := req["id"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "failed to get ration ID from the request")
		}

		rationAsset, err := rationKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get ration asset from the ledger")
		}

		// Update the ration asset with the provided information
		rationMap := (map[string]interface{})(*rationAsset)
		if category, ok := req["category"].(string); ok {
			rationMap["category"] = category
		}
		if description, ok := req["description"].(string); ok {
			rationMap["description"] = description
		}
		if rationPackage, ok := req["package"].(string); ok {
			rationMap["package"] = rationPackage
		}
		if distributedBy, ok := req["distributedBy"].(string); ok {
			rationMap["distributedBy"] = distributedBy
		}
		if quantity, ok := req["quantity"].(int); ok {
			rationMap["quantity"] = quantity
		}
		if expiryDate, ok := req["expiryDate"].(time.Time); ok {
			rationMap["expiryDate"] = expiryDate
		}
		if mfgDate, ok := req["mfgDate"].(time.Time); ok {
			rationMap["mfgDate"] = mfgDate
		}
		if batchNumber, ok := req["batchNumber"].(int); ok {
			rationMap["batchNumber"] = batchNumber
		}

		updatedRationAsset, err := rationAsset.Update(stub, rationMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update ration asset")
		}

		// Marshal asset back to JSON format
		updatedRationJSON, nerr := json.Marshal(updatedRationAsset)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		// Marshal message to be logged
		logMsg, erre := json.Marshal(fmt.Sprintf("Ration updated: %s", id))
		if erre != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		events.CallEvent(stub, "rationUpdatedLog", logMsg)

		return updatedRationJSON, nil
	},
}
