package datatypes

import (
	"encoding/json"
	"regexp"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type ContactInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var contactInfo = assets.DataType{
	AcceptedFormats: []string{"@object"},
	Description:     "A JSON string representing contact information with fields 'name', 'email', and 'phone'.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		dataStr, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a JSON string", 400)
		}

		var contact ContactInfo
		err := json.Unmarshal([]byte(dataStr), &contact)
		if err != nil {
			return "", nil, errors.WrapErrorWithStatus(err, "invalid JSON format", 400)
		}

		if contact.Name == "" {
			return "", nil, errors.NewCCError("name is required", 400)
		}

		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(contact.Email) {
			return "", nil, errors.NewCCError("invalid email format", 400)
		}

		phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
		if !phoneRegex.MatchString(contact.Phone) {
			return "", nil, errors.NewCCError("invalid phone format", 400)
		}

		return dataStr, contact, nil
	},
}
