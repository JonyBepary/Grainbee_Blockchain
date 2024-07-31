package datatypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// CustomDataTypes contain the user-defined primary data types
var CustomDataTypes = map[string]assets.DataType{
	"coordinates":               coordinates,
	"contactInfo":               contactInfo,
	"rationPickupSchedule":      rationPickupSchedule,
	"packageType":               packageType,
	"nid":                       nid,
	"rationDistributionHistory": rationDistributionHistory,
	"rationCardCategory":        rationCardCategory,
	"address":                   address,
	"rationCardStatus":          rationCardStatus,
	"operatingHours":            operatingHours,
	"inspectionStatus":          inspectionStatus,
	"rationTransaction":         rationTransaction,
	"rationCategory":            rationCategory,
}
