package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var ReadTotalRationsByDistributionPoint = tx.Transaction{
	Tag:         "readTotalRationsByDistributionPoint",
	Label:       "Read Total Rations by Distribution Point",
	Description: "Read the total number of rations by a distribution point",
	Method:      "GET",
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
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		distributionPointId, _ := req["distributionPointId"].(string)

		// Construct the CouchDB query
		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"@assetType":          "ration",
				"distributionPointId": distributionPointId,
			},
		}

		var err error
		response, err := assets.Search(stub, query, "", true)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error searching for book's author", 500)
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error marshaling response", 500)
		}

		return responseJSON, nil
	},
}
