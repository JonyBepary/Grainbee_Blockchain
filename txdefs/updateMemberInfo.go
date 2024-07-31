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

var UpdateMemberInfo = tx.Transaction{
	Tag:         "updateMemberInfo",
	Label:       "Update Member Information",
	Description: "Update the information of a member",
	Method:      "PUT",
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
			Tag:         "nid",
			Label:       "Member NID",
			Description: "Member NID",
			DataType:    "nid",
			Required:    true,
		},
		{
			Tag:         "name",
			Label:       "Member Name",
			Description: "Member Name",
			DataType:    "string",
			Required:    false,
		},
		{
			Tag:         "dateOfBirth",
			Label:       "Date of Birth",
			Description: "Date of Birth",
			DataType:    "datetime",
			Required:    false,
		},
		{
			Tag:         "height",
			Label:       "Member's Height",
			Description: "Member's Height",
			DataType:    "number",
			Required:    false,
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
			Tag:         "familySize",
			Label:       "Family Size",
			Description: "Family Size",
			DataType:    "integer",
			Required:    false,
		},
		{
			Tag:         "income",
			Label:       "Income",
			Description: "Income",
			DataType:    "integer",
			Required:    false,
		},
		{
			Tag:         "disabilityStatus",
			Label:       "Disability Status",
			Description: "Disability Status",
			DataType:    "boolean",
			Required:    false,
		},
		{
			Tag:         "rationCardNumber",
			Label:       "Ration Card Number",
			Description: "Ration Card Number",
			DataType:    "string",
			Required:    false,
		},
		{
			Tag:         "rationCardStatus",
			Label:       "Ration Card Status",
			Description: "Ration Card Status",
			DataType:    "rationCardStatus",
			Required:    false,
		},
		{
			Tag:         "rationCardIssuedDate",
			Label:       "Ration Card Issued Date",
			Description: "Ration Card Issued Date",
			DataType:    "datetime",
			Required:    false,
		},
		{
			Tag:         "rationCardExpiryDate",
			Label:       "Ration Card Expiry Date",
			Description: "Ration Card Expiry Date",
			DataType:    "datetime",
			Required:    false,
		},
		{
			Tag:         "rationCardCategory",
			Label:       "Ration Card Category",
			Description: "Ration Card Category",
			DataType:    "rationCardCategory",
			Required:    false,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		nid, _ := req["nid"].(string)

		// Retrieve the member asset
		memberKey, ok := req["member"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter member must be an asset")
		}
		memberAsset, err := memberKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get member asset from the ledger")
		}

		// Update the member asset with the provided information
		memberMap := (map[string]interface{})(*memberAsset)
		if name, ok := req["name"].(string); ok {
			memberMap["name"] = name
		}
		if dateOfBirth, ok := req["dateOfBirth"].(string); ok {
			memberMap["dateOfBirth"] = dateOfBirth
		}
		if height, ok := req["height"].(float64); ok {
			memberMap["height"] = height
		}
		if address, ok := req["address"].(string); ok {
			memberMap["address"] = address
		}
		if contactInformation, ok := req["contactInformation"].(string); ok {
			memberMap["contactInformation"] = contactInformation
		}
		if familySize, ok := req["familySize"].(int); ok {
			memberMap["familySize"] = familySize
		}
		if income, ok := req["income"].(int); ok {
			memberMap["income"] = income
		}
		if disabilityStatus, ok := req["disabilityStatus"].(bool); ok {
			memberMap["disabilityStatus"] = disabilityStatus
		}
		if rationCardNumber, ok := req["rationCardNumber"].(string); ok {
			memberMap["rationCardNumber"] = rationCardNumber
		}
		if rationCardStatus, ok := req["rationCardStatus"].(string); ok {
			memberMap["rationCardStatus"] = rationCardStatus
		}
		if rationCardIssuedDate, ok := req["rationCardIssuedDate"].(string); ok {
			memberMap["rationCardIssuedDate"] = rationCardIssuedDate
		}
		if rationCardExpiryDate, ok := req["rationCardExpiryDate"].(string); ok {
			memberMap["rationCardExpiryDate"] = rationCardExpiryDate
		}
		if rationCardCategory, ok := req["rationCardCategory"].(string); ok {
			memberMap["rationCardCategory"] = rationCardCategory
		}

		updatedMemberAsset, err := memberAsset.Update(stub, memberMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update member asset")
		}

		// Marshal asset back to JSON format
		updatedMemberJSON, nerr := json.Marshal(updatedMemberAsset)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		// Marshal message to be logged
		logMsg, erre := json.Marshal(fmt.Sprintf("Member information updated for NID: %s", nid))
		if erre != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		events.CallEvent(stub, "memberInfoUpdatedLog", logMsg)

		return updatedMemberJSON, nil
	},
}
