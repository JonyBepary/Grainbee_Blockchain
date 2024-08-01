package datatypes

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

// Example of a custom data type using enum-like structure (iota)
// This allows the use of verification by const values instead of float64, improving readability
// Example:

type PackageType float64

const (
	PackageTypePacket PackageType = iota
	PackageTypeBottle
	PackageTypeCan
	PackageTypeBox
	PackageTypeSachet
	PackageTypeBag
	PackageTypeRamadan
)

// CheckType checks if the given value is defined as valid PackageType consts
func (b PackageType) CheckType() errors.ICCError {
	switch b {
	case PackageTypePacket:
		return nil
	case PackageTypeBottle:
		return nil
	case PackageTypeCan:
		return nil
	case PackageTypeBox:
		return nil
	case PackageTypeSachet:
		return nil
	case PackageTypeBag:
		return nil
	case PackageTypeRamadan:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}

}

var packageType = assets.DataType{
	AcceptedFormats: []string{"number"},
	DropDownValues: map[string]interface{}{
		"Packet":  PackageTypePacket,
		"Bottle":  PackageTypeBottle,
		"Can":     PackageTypeCan,
		"Box":     PackageTypeBox,
		"Sachet":  PackageTypeSachet,
		"Bag":     PackageTypeBag,
		"Ramadan": PackageTypeRamadan,
	},
	Description: ``,

	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal float64
		switch v := data.(type) {
		case float64:
			dataVal = v
		case int:
			dataVal = (float64)(v)
		case PackageType:
			dataVal = (float64)(v)
		case string:
			var err error
			dataVal, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "asset property must be an integer, is %t", 400)
			}
		default:
			return "", nil, errors.NewCCError("asset property must be an integer, is %t", 400)
		}
		// for ramanadan package type we need to check if the date is ramadan
		// if it is ramadan then only we can use this package type
		if dataVal == (float64)(PackageTypeRamadan) {
			date := time.Now()
			if date.Month() != 4 {
				return "", nil, errors.NewCCError("Ramadan package type can only be used in the month of Ramadan ", 400)
			}
		}

		retVal := (PackageType)(dataVal)
		err := retVal.CheckType()
		return fmt.Sprint(retVal), retVal, err
	},
}
