package datatypes

import (
	"regexp"
	"strings"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

var nid = assets.DataType{
	AcceptedFormats: []string{"string"},
	Description:     "A string representing a Bangladeshi NID number.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		nidStr, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}

		// Remove any non-digit characters
		nidStr = strings.ReplaceAll(nidStr, "-", "")
		nidStr = strings.ReplaceAll(nidStr, " ", "")

		// Check if the NID number is either 10, 13 or 17 digits long
		nidRegex := regexp.MustCompile(`^\d{10}$|^\d{13}$|^\d{17}$`)
		if !nidRegex.MatchString(nidStr) {
			return "", nil, errors.NewCCError("invalid NID number", 400)
		}
		// first number must be 1-9
		if nidStr[0] == '0' {
			return "", nil, errors.NewCCError("invalid NID number", 400)
		}
		// if 17 digits, first 2 digits must be 19 or 20 regex
		if len(nidStr) == 17 {
			nidRegex = regexp.MustCompile(`^(19|20)\d{15}$`)
			if !nidRegex.MatchString(nidStr) {
				return "", nil, errors.NewCCError("invalid NID number", 400)
			}
		}
		// if len is 13 then return error asking birth year at the front
		if len(nidStr) == 13 {
			return "", nil, errors.NewCCError("invalid NID number, please add birth year at the front", 400)
		}
		return nidStr, nidStr, nil
	},
}
