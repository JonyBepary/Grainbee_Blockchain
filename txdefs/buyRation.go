package txdefs

import (
	"encoding/json"
	"sync"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var BuyRation = tx.Transaction{
	Tag:         "buyRation",
	Label:       "Buy Ration",
	Description: "Customer buys a ration",
	Method:      "POST",
	Callers: []accesscontrol.Caller{
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
			Tag:         "rationCardNumber",
			Label:       "Ration Card Number",
			Description: "Ration Card Number",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "rationId",
			Label:       "Ration ID",
			Description: "Ration ID",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "quantity",
			Label:       "Quantity",
			Description: "Quantity of rations to buy",
			DataType:    "integer",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		customerKey, ok := req["customer"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter customer must be an asset")
		}
		newCustomerKey, ok := req["newCustomer"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter newCustomer must be an asset")
		}

		// Returns customer from channel
		customerMap, err := customerKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "failed to get asset from the ledger", err.Status())
		}

		// Returns new customer from channel
		newCustomerMap, err := newCustomerKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "failed to get asset from the ledger", err.Status())
		}

		// Use a WaitGroup to wait for all updates to complete
		var wg sync.WaitGroup
		errChan := make(chan error, 1)

		updateField := func(field string, value interface{}) {
			defer wg.Done()
			if value != nil {
				customerMap[field] = value
			}
		}

		fields := map[string]interface{}{
			"name":    newCustomerMap["name"],
			"address": newCustomerMap["address"],
			"phone":   newCustomerMap["phone"],
			"email":   newCustomerMap["email"],
		}

		for field, value := range fields {
			wg.Add(1)
			go updateField(field, value)
		}

		wg.Wait()
		close(errChan)

		for err := range errChan {
			if err != nil {
				return nil, errors.WrapError(err, "failed to update customer asset")
			}
		}

		// Update customer in the ledger
		updatedCustomerMap, err := customerKey.Update(stub, customerMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update asset")
		}

		// Marshal updated customer map to JSON
		updatedCustomerJSON, nerr := json.Marshal(updatedCustomerMap)
		if nerr != nil {
			return nil, errors.WrapError(nerr, "failed to marshal response")
		}

		return updatedCustomerJSON, nil
	},
}
