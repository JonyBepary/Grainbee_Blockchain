package datatypes

import (
	"fmt"
	"strconv"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type InspectionStatus float64

const (
	InspectionStatusPending InspectionStatus = iota
	InspectionStatusInProgress
	InspectionStatusCompleted
	InspectionStatusFailed
	InspectionStatusCancelled
)

// CheckType checks if the given value is defined as valid InspectionStatus consts
func (i InspectionStatus) CheckType() errors.ICCError {
	switch i {
	case InspectionStatusPending, InspectionStatusInProgress, InspectionStatusCompleted, InspectionStatusFailed, InspectionStatusCancelled:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}
}

var inspectionStatus = assets.DataType{
	AcceptedFormats: []string{"string"},
	DropDownValues: map[string]interface{}{
		"Pending":    InspectionStatusPending,
		"InProgress": InspectionStatusInProgress,
		"Completed":  InspectionStatusCompleted,
		"Failed":     InspectionStatusFailed,
		"Cancelled":  InspectionStatusCancelled,
	},
	Description: "A string representing the inspection status.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal float64
		switch v := data.(type) {
		case float64:
			dataVal = v
		case int:
			dataVal = float64(v)
		case InspectionStatus:
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

		retVal := (InspectionStatus)(dataVal)
		err := retVal.CheckType()
		return fmt.Sprint(retVal), retVal, err
	},
}
