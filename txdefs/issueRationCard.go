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

var IssueRationCard = tx.Transaction{
	Tag:         "issueRationCard",
	Label:       "Issue Ration Card",
	Description: "Issue a ration card for a member",
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
			Tag:         "nid",
			Label:       "Member NID",
			Description: "Member NID",
			DataType:    "nid",
			Required:    true,
		},
		{
			Tag:         "rationCardNumber",
			Label:       "Ration Card Number",
			Description: "Ration Card Number",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "rationCardStatus",
			Label:       "Ration Card Status",
			Description: "Ration Card Status",
			DataType:    "rationCardStatus",
			Required:    true,
		},
		{
			Tag:         "rationCardIssuedDate",
			Label:       "Ration Card Issued Date",
			Description: "Ration Card Issued Date",
			DataType:    "datetime",
			Required:    true,
		},
		{
			Tag:         "rationCardExpiryDate",
			Label:       "Ration Card Expiry Date",
			Description: "Ration Card Expiry Date",
			DataType:    "datetime",
			Required:    true,
		},
		{
			Tag:         "rationCardCategory",
			Label:       "Ration Card Category",
			Description: "Ration Card Category",
			DataType:    "rationCardCategory",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		nid, _ := req["nid"].(string)
		rationCardNumber, _ := req["rationCardNumber"].(string)
		rationCardStatus, _ := req["rationCardStatus"].(string)
		rationCardIssuedDate, _ := req["rationCardIssuedDate"].(string)
		rationCardExpiryDate, _ := req["rationCardExpiryDate"].(string)
		rationCardCategory, _ := req["rationCardCategory"].(string)

		MemberKey, ok := req["member"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter member must be an asset")
		}

		memberAsset, err := MemberKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get member asset from the ledger")
		}

		// Check if ration card already issued
		if memberAsset != nil {
			return nil, errors.WrapError(nil, "Ration card already issued for this member")
		}

		// Update the member asset with ration card details
		memberMap := (map[string]interface{})(*memberAsset)
		memberMap["rationCardNumber"] = rationCardNumber
		memberMap["rationCardStatus"] = rationCardStatus
		memberMap["rationCardIssuedDate"] = rationCardIssuedDate
		memberMap["rationCardExpiryDate"] = rationCardExpiryDate
		memberMap["rationCardCategory"] = rationCardCategory

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
		logMsg, erre := json.Marshal(fmt.Sprintf("Ration card issued for member with NID: %s", nid))
		if erre != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		events.CallEvent(stub, "rationCardIssuedLog", logMsg)

		return updatedMemberJSON, nil
	},
}
