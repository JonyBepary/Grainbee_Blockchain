// [
//   {
//     "transactionID": "T123456789",
//     "transactionType": "distributorReceive",
//     "itemName": "Rice",
//     "quantity": 50,
//     "unit": "kg",
//     "transactionDate": "2023-10-05T14:30:00Z",
//     "distributorID": "D123456789",
//     "source": "Government Agency"
//   },
//   {
//     "transactionID": "T987654321",
//     "transactionType": "memberReceive",
//     "itemName": "Rice",
//     "quantity": 10,
//     "unit": "kg",
//     "transactionDate": "2023-10-06T10:00:00Z",
//     "distributorID": "D123456789",
//     "memberID": "M123456789"
//   }
// ]

package datatypes

import (
	"encoding/json"
	"time"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type RationTransaction struct {
	TransactionID   string `json:"transactionID"`
	TransactionType string `json:"transactionType"`
	ItemName        string `json:"itemName"`
	Quantity        int    `json:"quantity"`
	Unit            string `json:"unit"`
	TransactionDate string `json:"transactionDate"`
	DistributorID   string `json:"distributorID"`
	Source          string `json:"source,omitempty"`
	MemberID        string `json:"memberID,omitempty"`
}

var rationTransaction = assets.DataType{
	AcceptedFormats: []string{"@object"},
	Description:     "A JSON array representing ration transactions with fields for both distributor and member transactions.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		dataStr, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a JSON string", 400)
		}

		var transactions []RationTransaction
		err := json.Unmarshal([]byte(dataStr), &transactions)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid JSON format", 400)
		}

		validTransactionTypes := map[string]bool{
			"distributorReceive": true,
			"memberReceive":      true,
		}

		for _, transaction := range transactions {
			if transaction.TransactionID == "" {
				return "", nil, errors.NewCCError("transactionID is required", 400)
			}

			if !validTransactionTypes[transaction.TransactionType] {
				return "", nil, errors.NewCCError("invalid transactionType", 400)
			}

			if transaction.ItemName == "" {
				return "", nil, errors.NewCCError("itemName is required", 400)
			}

			if transaction.Quantity <= 0 {
				return "", nil, errors.NewCCError("quantity must be greater than 0", 400)
			}

			if transaction.Unit == "" {
				return "", nil, errors.NewCCError("unit is required", 400)
			}

			if transaction.TransactionDate == "" {
				return "", nil, errors.NewCCError("transactionDate is required", 400)
			}

			_, err := time.Parse(time.RFC3339, transaction.TransactionDate)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "invalid transactionDate format", 400)
			}

			if transaction.DistributorID == "" {
				return "", nil, errors.NewCCError("distributorID is required", 400)
			}

			if transaction.TransactionType == "distributorReceive" && transaction.Source == "" {
				return "", nil, errors.NewCCError("source is required for distributorReceive transactions", 400)
			}

			if transaction.TransactionType == "memberReceive" && transaction.MemberID == "" {
				return "", nil, errors.NewCCError("memberID is required for memberReceive transactions", 400)
			}
		}

		return dataStr, transactions, nil
	},
}
