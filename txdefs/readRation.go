package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var ReadRation = tx.Transaction{
	Tag:         "readRation",
	Label:       "Read Ration",
	Description: "Read the information of a ration",
	Method:      "GET",
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
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {

		name, ok := req["name"].(string)
		if !ok {
			return nil, errors.NewCCError("name is required", 400)
		}

		// Prepare couchdb query
		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"@assetType": "ration",
				"name":       name,
			},
		}

		// Retrieve the ration asset
		rationAsset, err := assets.Search(stub, query, "", true)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get ration asset from the ledger")
		}

		// Marshal asset back to JSON format
		rationJSON, nerr := json.Marshal(rationAsset)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return rationJSON, nil
	},
}
