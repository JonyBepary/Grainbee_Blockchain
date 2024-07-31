package datatypes

import (
	"encoding/json"
	"regexp"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

var address = assets.DataType{
	AcceptedFormats: []string{"@object"},
	Description:     "A JSON string representing an address with fields 'street', 'city', 'state', 'postalCode', and 'country'.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		dataStr, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a JSON string", 400)
		}

		var addr Address
		err := json.Unmarshal([]byte(dataStr), &addr)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid JSON format", 400)
		}

		if addr.Street == "" {
			return "", nil, errors.NewCCError("street is required", 400)
		}

		if addr.City == "" {
			return "", nil, errors.NewCCError("city is required", 400)
		}

		if addr.State == "" {
			return "", nil, errors.NewCCError("state is required", 400)
		}

		postalCodeRegex := regexp.MustCompile(`^\d{5}(-\d{4})?$`)
		if !postalCodeRegex.MatchString(addr.PostalCode) {
			return "", nil, errors.NewCCError("invalid postalCode format", 400)
		}

		if addr.Country == "" {
			return "", nil, errors.NewCCError("country is required", 400)
		}

		return dataStr, addr, nil
	},
}
