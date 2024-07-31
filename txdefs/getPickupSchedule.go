package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var GetPickupSchedule = tx.Transaction{
	Tag:         "getPickupSchedule",
	Label:       "Get Pickup Schedule",
	Description: "Get the pickup schedule for a specific distribution point",
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

		// Retrieve the distribution point asset
		distributionPointKey, ok := req[distributionPointId].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "failed to get distribution point ID from the request")
		}

		distributionPointAsset, err := distributionPointKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get distribution point asset from the ledger")
		}

		// Get the pickup schedule from the distribution point asset
		pickupSchedule, ok := distributionPointAsset.GetProp("pickupSchedule").(map[string]interface{})
		if !ok {
			return nil, errors.NewCCError("pickupSchedule property not found", 404)
		}

		// Marshal the pickup schedule back to JSON format
		pickupScheduleJSON, nerr := json.Marshal(pickupSchedule)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return pickupScheduleJSON, nil
	},
}
