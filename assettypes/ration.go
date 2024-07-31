package assettypes

import (
	"time"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

const (
	MinRationLimit = 1
	MaxRationLimit = 100000
	MinExpiryDate  = 60
)

// type Ration struct {
// 	ID             string    `json:"id"`
// 	Type           string    `json:"type"`
// 	Quantity       int       `json:"quantity"`
// 	ExpiryDate     time.Time `json:"expiryDate"`
// 	ManufacturingDate time.Time `json:"mfgDate"`
// 	BatchNumber    int       `json:"batchNumber"`
// 	Description    string    `json:"description"`
// 	Package        string    `json:"package"`
// 	DistributedBy  string    `json:"distributedBy"`
// }

// isValidRationType checks if the given ration type is valid

func isValidRationPackage(p string) bool {
	// hashmap of valid ration packages
	rationPackages := map[string]bool{
		"packet": true,
		"bottle": true,
		"can":    true,
		"box":    true,
		"sachet": true,
		"bag":    true,
	}
	if _, ok := rationPackages[p]; ok {
		return true
	}
	return false
}

// Description of a Ration
var RationAsset = assets.AssetType{
	Tag:         "ration",
	Label:       "Ration",
	Description: "Ration",

	Props: []assets.AssetProp{
		//id
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "Ration ID",
			DataType: "string",                      // Datatypes are identified at datatypes folder
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)

		},
		//type
		{
			// Composite Key
			Required: true,
			Tag:      "category",
			Label:    "Ration Category",
			DataType: "rationCategory",              // values: can be food, water, medicine, etc.
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)

		},
		// Description
		{
			// package description
			Required: false,
			Tag:      "description",
			Label:    "Ration Description",
			DataType: "string",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
		},

		// Package
		{
			// package
			Required:     true,
			Tag:          "package",
			Label:        "Ration Package",
			DataType:     "string",                      // values: are packet, bottle, can, box, sachet, bag.
			Writers:      []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
			DefaultValue: "packet",
			Validate: func(rationPackage interface{}) error {
				if !isValidRationPackage(rationPackage.(string)) {
					return errors.NewCCError("Invalid ration package", 400)
				}
				return nil
			},
		},
		//Distribution
		{
			// Composite Key
			Required: true,
			Tag:      "distributedBy",
			Label:    "distributedBy",
			DataType: "->distributor",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
			Validate: func(sellerId interface{}) error {
				if len(sellerId.(string)) < 16 {
					return errors.NewCCError("Seller ID must be at least 16 characters", 400)
				}
				// check if sellerID is valid

				return nil
			},
		},
		//Quantity
		{
			// Composite Key
			Required:     true,
			IsKey:        true,
			Tag:          "quantity",
			Label:        "Ration Quantity",
			DataType:     "int",
			Writers:      []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
			DefaultValue: 1,
			Validate: func(quantity interface{}) error {
				if quantity.(int) < MinRationLimit || quantity.(int) > MaxRationLimit {
					return errors.NewCCError("Quantity must be between 1 and 100000", 400)
				}
				return nil
			},
		},
		// EXPIRY DATE
		{
			// Composite Key
			Required: true,
			Tag:      "expiryDate",
			Label:    "Expiry Date",
			DataType: "datetime",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
			Validate: func(expiryDate interface{}) error {
				// check if expiry date is valid (not in the past and not more than 60 DAYS in the future)
				if expiryDate.(time.Time).Before(time.Now()) {
					return errors.NewCCError("Expiry date must be in the future", 400)
				}
				if expiryDate.(time.Time).After(time.Now().AddDate(0, 0, MinExpiryDate)) {
					return errors.NewCCError("Expiry date must be within 60 days", 400)
				}
				return nil
			},
		},
		// mfg date
		{
			// Composite Key
			Required: true,
			Tag:      "mfgDate",
			Label:    "Manufacturing Date",
			DataType: "datetime",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
			Validate: func(mfgDate interface{}) error {
				// check if mfg date is valid (not in the future)
				if mfgDate.(time.Time).After(time.Now()) {
					return errors.NewCCError("Manufacturing date must be in the past", 400)
				}
				return nil
			},
		},
		// BATCH NUMBER
		{
			// Composite Key
			Required: true,
			Tag:      "batchNumber",
			Label:    "Batch Number",
			DataType: "int",
			Writers:  []string{`org2MSP`, "orgMSP"}, // This means only org2 can create the asset (others can edit)
			Validate: func(batchNumber interface{}) error {
				// check if batch number is valid
				if batchNumber.(int) < 1 {
					return errors.NewCCError("Batch number must be greater than 0", 400)
				}
				return nil
			},
		},
	},
}
