package datatypes

import (
	"encoding/json"
	"time"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

/*
	{
	  "days": ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"],
	  "openingTime": "2023-10-05T09:00:00Z",
	  "closingTime": "2023-10-05T17:00:00Z"
	}
*/
type OperatingHours struct {
	Days        []string `json:"days"`
	OpeningTime string   `json:"openingTime"`
	ClosingTime string   `json:"closingTime"`
}

var operatingHours = assets.DataType{
	AcceptedFormats: []string{"@object"},
	Description:     "A JSON string representing the operating hours of a distribution point with fields 'days', 'openingTime', and 'closingTime'.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		dataStr, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a JSON string", 400)
		}

		var hours OperatingHours
		err := json.Unmarshal([]byte(dataStr), &hours)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid JSON format", 400)
		}

		if len(hours.Days) == 0 {
			return "", nil, errors.NewCCError("days are required", 400)
		}

		validDays := map[string]bool{
			"Monday":    true,
			"Tuesday":   true,
			"Wednesday": true,
			"Thursday":  true,
			"Friday":    true,
			"Saturday":  true,
			"Sunday":    true,
		}

		for _, day := range hours.Days {
			if !validDays[day] {
				return "", nil, errors.NewCCError("invalid day", 400)
			}
		}

		openingTime, err := time.Parse(time.RFC3339, hours.OpeningTime)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid openingTime format", 400)
		}

		closingTime, err := time.Parse(time.RFC3339, hours.ClosingTime)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid closingTime format", 400)
		}

		if closingTime.Before(openingTime) {
			return "", nil, errors.NewCCError("closingTime must be after openingTime", 400)
		}

		return dataStr, hours, nil
	},
}
