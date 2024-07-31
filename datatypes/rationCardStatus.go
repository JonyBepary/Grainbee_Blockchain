package datatypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type RationCardStatus string

const (
	RationCardStatusActive    RationCardStatus = "active"
	RationCardStatusInactive  RationCardStatus = "inactive"
	RationCardStatusSuspended RationCardStatus = "suspended"
	RationCardStatusExpired   RationCardStatus = "expired"
	RationCardStatusPending   RationCardStatus = "pending"
)

var rationCardStatus = assets.DataType{
	AcceptedFormats: []string{"string"},
	DropDownValues: map[string]interface{}{
		"Active":    RationCardStatusActive,
		"Inactive":  RationCardStatusInactive,
		"Suspended": RationCardStatusSuspended,
		"Expired":   RationCardStatusExpired,
		"Pending":   RationCardStatusPending,
	},
	Description: "A string representing the ration card status.",
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal string
		switch v := data.(type) {
		case string:
			dataVal = v
		default:
			return "", nil, errors.NewCCError("invalid type", 400)
		}

		status := RationCardStatus(dataVal)
		statusStr := string(status)
		switch status {
		case RationCardStatusActive, RationCardStatusInactive, RationCardStatusSuspended, RationCardStatusExpired, RationCardStatusPending:
			break
		default:
			return "", nil, errors.NewCCError("invalid type", 400)
		}
		return statusStr, status, nil
	},
}
