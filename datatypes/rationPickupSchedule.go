package datatypes

import (
	"encoding/json"
	"time"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type RationPickupSchedule struct {
	PickupDate string `json:"pickupDate"`
	Location   string `json:"location"`
	RationType string `json:"rationType"`
	Quantity   int    `json:"quantity"`
}

var rationPickupSchedule = assets.DataType{
	AcceptedFormats: []string{"@object"},
	Description:     "A JSON string representing a ration pickup schedule with fields 'pickupDate', 'location', 'rationType', and 'quantity'.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		dataStr, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a JSON string", 400)
		}

		var schedule RationPickupSchedule
		err := json.Unmarshal([]byte(dataStr), &schedule)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid JSON format", 400)
		}

		if schedule.PickupDate == "" {
			return "", nil, errors.NewCCError("pickupDate is required", 400)
		}

		_, err = time.Parse(time.RFC3339, schedule.PickupDate)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid pickupDate format", 400)
		}

		if schedule.Location == "" {
			return "", nil, errors.NewCCError("location is required", 400)
		}

		if schedule.RationType == "" {
			return "", nil, errors.NewCCError("rationType is required", 400)
		}

		if schedule.Quantity <= 0 {
			return "", nil, errors.NewCCError("quantity must be greater than 0", 400)
		}

		return dataStr, schedule, nil
	},
}
