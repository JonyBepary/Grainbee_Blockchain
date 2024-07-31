package datatypes

import (
	"encoding/json"
	"time"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type RationDistributionHistory struct {
	DistributionID   string `json:"distributionID"`
	DistributionDate string `json:"distributionDate"`
	RationType       string `json:"rationType"`
	Quantity         int    `json:"quantity"`
	DistributedTo    string `json:"distributedTo"`
	Location         string `json:"location"`
}

var rationDistributionHistory = assets.DataType{
	AcceptedFormats: []string{"@object"},
	Description:     "A JSON string representing ration distribution history with fields 'distributionID', 'distributionDate', 'rationType', 'quantity', 'distributedTo', and 'location'.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		dataStr, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a JSON string", 400)
		}

		var history RationDistributionHistory
		err := json.Unmarshal([]byte(dataStr), &history)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid JSON format", 400)
		}

		if history.DistributionID == "" {
			return "", nil, errors.NewCCError("distributionID is required", 400)
		}

		if history.DistributionDate == "" {
			return "", nil, errors.NewCCError("distributionDate is required", 400)
		}

		_, err = time.Parse(time.RFC3339, history.DistributionDate)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid distributionDate format", 400)
		}

		if history.RationType == "" {
			return "", nil, errors.NewCCError("rationType is required", 400)
		}

		if history.Quantity <= 0 {
			return "", nil, errors.NewCCError("quantity must be greater than 0", 400)
		}

		if history.DistributedTo == "" {
			return "", nil, errors.NewCCError("distributedTo is required", 400)
		}

		if history.Location == "" {
			return "", nil, errors.NewCCError("location is required", 400)
		}

		return dataStr, history, nil
	},
}
