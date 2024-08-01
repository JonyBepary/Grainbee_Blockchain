package datatypes

import (
	"fmt"
	"strconv"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type RationCardCategory float64

const (
	RationCardCategorySingle RationCardCategory = iota
	RationCardCategorySmall
	RationCardCategoryMedium
	RationCardCategoryLarge
)

// CheckType checks if the given value is defined as valid RationCardCategory consts
func (r RationCardCategory) CheckType() errors.ICCError {
	switch r {
	case RationCardCategorySingle, RationCardCategorySmall, RationCardCategoryMedium, RationCardCategoryLarge:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}
}

var rationCardCategory = assets.DataType{
	AcceptedFormats: []string{"string"},
	DropDownValues: map[string]interface{}{
		"Single": RationCardCategorySingle,
		"Small":  RationCardCategorySmall,
		"Medium": RationCardCategoryMedium,
		"Large":  RationCardCategoryLarge,
	},
	Description: "A string representing the ration card category.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal float64
		switch v := data.(type) {
		case float64:
			dataVal = v
		case int:
			dataVal = (float64)(v)
		case RationCardCategory:
			dataVal = (float64)(v)
		case string:
			var err error
			dataVal, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "invalid number format", 400)
			}
		default:
			return "", nil, errors.NewCCError("invalid type", 400)
		}

		retVal := (RationCardCategory)(dataVal)
		err := retVal.CheckType()
		return fmt.Sprint(retVal), retVal, err
	},
}
