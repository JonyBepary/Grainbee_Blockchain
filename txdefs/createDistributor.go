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

var CreateDistributor = tx.Transaction{
	Tag:         "createDistributor",
	Label:       "Create Distributor",
	Description: "Create a new distributor",
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
			Tag:         "distributorId",
			Label:       "Distributor ID",
			Description: "Distributor ID",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "name",
			Label:       "Name of the distributor",
			Description: "Name of the distributor",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "address",
			Label:       "Address",
			Description: "Address",
			DataType:    "address",
			Required:    false,
		},
		{
			Tag:         "contactInformation",
			Label:       "Contact Information",
			Description: "Contact Information",
			DataType:    "contactInfo",
			Required:    false,
		},
		{
			Tag:         "licenseNumber",
			Label:       "License Number",
			Description: "License Number",
			DataType:    "string",
			Required:    false,
		},
		{
			Tag:         "licenseIssueDate",
			Label:       "License Issue Date",
			Description: "License Issue Date",
			DataType:    "datetime",
			Required:    false,
		},
		{
			Tag:         "licenseExpiryDate",
			Label:       "License Expiry Date",
			Description: "License Expiry Date",
			DataType:    "datetime",
			Required:    false,
		},
		{
			Tag:         "distributionArea",
			Label:       "Distribution Area",
			Description: "Distribution Area",
			DataType:    "string",
			Required:    false,
		},
		{
			Tag:         "lastInspectionDate",
			Label:       "Last Inspection Date",
			Description: "Last Inspection Date",
			DataType:    "datetime",
			Required:    false,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		distributorId, _ := req["distributorId"].(string)
		name, _ := req["name"].(string)
		address, _ := req["address"].(string)
		contactInformation, _ := req["contactInformation"].(string)
		licenseNumber, _ := req["licenseNumber"].(string)
		licenseIssueDate, _ := req["licenseIssueDate"].(time.Time)
		licenseExpiryDate, _ := req["licenseExpiryDate"].(time.Time)
		distributionArea, _ := req["distributionArea"].(string)
		lastInspectionDate, _ := req["lastInspectionDate"].(time.Time)

		distributorMap := make(map[string]interface{})
		distributorMap["@assetType"] = "distributor"
		distributorMap["distributorId"] = distributorId
		distributorMap["name"] = name
		distributorMap["address"] = address
		distributorMap["contactInformation"] = contactInformation
		distributorMap["licenseNumber"] = licenseNumber
		distributorMap["licenseIssueDate"] = licenseIssueDate
		distributorMap["licenseExpiryDate"] = licenseExpiryDate
		distributorMap["distributionArea"] = distributionArea
		distributorMap["lastInspectionDate"] = lastInspectionDate

		distributorAsset, err := assets.NewAsset(distributorMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = distributorAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		distributorJSON, nerr := json.Marshal(distributorAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		logMsg, ok := json.Marshal(fmt.Sprintf("New distributor created: %s", distributorId))
		if ok != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		events.CallEvent(stub, "distributorCreatedLog", logMsg)

		return distributorJSON, nil
	},
}
