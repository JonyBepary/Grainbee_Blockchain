package datatypes

import (
	"fmt"
	"strconv"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type RationCategory float64

const (
	RationCategoryGrains RationCategory = iota
	RationCategoryOil
	RationCategoryPulses
	RationCategorySugar
	RationCategorySpices
	RationCategoryRamadanEssentials
	RationCategoryPandemicEssentials
	RationCategoryOthers
)

// CheckType checks if the given value is defined as valid RationCategory consts
func (i RationCategory) CheckType() errors.ICCError {
	switch i {
	case RationCategoryGrains, RationCategoryOil, RationCategoryPulses, RationCategorySugar, RationCategorySpices, RationCategoryRamadanEssentials, RationCategoryPandemicEssentials, RationCategoryOthers:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}
}

var rationCategory = assets.DataType{
	AcceptedFormats: []string{"string"},
	DropDownValues: map[string]interface{}{
		"Grains":              RationCategoryGrains,
		"Oil":                 RationCategoryOil,
		"Pulses":              RationCategoryPulses,
		"Sugar":               RationCategorySugar,
		"Spices":              RationCategorySpices,
		"Ramadan Essentials":  RationCategoryRamadanEssentials,
		"Pandemic Essentials": RationCategoryPandemicEssentials,
		"Others":              RationCategoryOthers,
	},
	Description: "A string representing the category of the ration Ration.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal float64
		switch v := data.(type) {
		case float64:
			dataVal = v
		case int:
			dataVal = float64(v)
		case RationCategory:
			dataVal = float64(v)
		case string:
			var err error
			dataVal, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "invalid number format", 400)
			}
		default:
			return "", nil, errors.NewCCError("invalid type", 400)
		}

		retVal := (RationCategory)(dataVal)
		err := retVal.CheckType()
		return fmt.Sprint(retVal), retVal, err
	},
}
