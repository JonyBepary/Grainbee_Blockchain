package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

// Return the number of Book of a library
// GET method
var GetNumberOfBookFromLibrary = tx.Transaction{
	Tag:         "getNumberOfBookFromLibrary",
	Label:       "Get Number Of book From Library",
	Description: "Return the number of book of a library",
	Method:      "GET",
	Callers: []accesscontrol.Caller{ // Only org2 can call this transaction
		{MSP: "org2MSP"},
		{MSP: "orgMSP"},
	},

	Args: []tx.Argument{
		{
			Tag:         "ration",
			Label:       "Ration",
			Description: "Ration",
			DataType:    "->ration",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		rationKey, _ := req["ration"].(assets.Key)

		// Returns ration from channel
		rationMap, err := rationKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "failed to get asset from the ledger", err.Status())
		}

		numberOfRation := 0
		ration, ok := rationMap["ration"].([]interface{})
		if ok {
			numberOfRation = len(ration)
		}

		returnMap := make(map[string]interface{})
		returnMap["numberOfRation"] = numberOfRation

		// Marshal asset back to JSON format
		returnJSON, nerr := json.Marshal(returnMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return returnJSON, nil
	},
}
