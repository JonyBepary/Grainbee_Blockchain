package assettypes

import (
	"fmt"

	"github.com/hyperledger-labs/cc-tools/assets"
)

var Member = assets.AssetType{
	Tag:         "member",
	Label:       "Member",
	Description: "Pernsonal infortion of a member",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "nid",
			Label:    "nid (Bangladeshi National Identity)",
			DataType: "nid",                         // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Name of the member",
			DataType: "string",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
			// Validate function
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("name must be non-empty")
				}
				return nil
			},
		},
		{
			// Optional property
			Tag:      "dateOfBirth",
			Label:    "Date of Birth",
			DataType: "datetime",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Property with default value
			Tag:          "height",
			Label:        "member's height",
			DefaultValue: 174,
			DataType:     "number",
			Writers:      []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: Address
			Tag:      "address",
			Label:    "Address",
			DataType: "address",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: ContactInformation
			Tag:      "contactInformation",
			Label:    "Contact Information",
			DataType: "contactInfo",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: FamilySize
			Tag:      "familySize",
			Label:    "Family Size",
			DataType: "integer",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: Income
			Tag:      "income",
			Label:    "Income",
			DataType: "integer",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: DisabilityStatus
			Tag:      "disabilityStatus",
			Label:    "Disability Status",
			DataType: "boolean",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: RationCardNumber
			Tag:      "rationCardNumber",
			Label:    "Ration Card Number",
			DataType: "string",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: RationCardStatus
			Tag:      "rationCardStatus",
			Label:    "Ration Card Status",
			DataType: "rationCardStatus",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: RationCardIssuedDate
			Tag:      "rationCardIssuedDate",
			Label:    "Ration Card Issued Date",
			DataType: "datetime",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: RationCardExpiryDate
			Tag:      "rationCardExpiryDate",
			Label:    "Ration Card Expiry Date",
			DataType: "datetime",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: RationCardCategory
			Tag:      "rationCardCategory",
			Label:    "Ration Card Category",
			DataType: "rationCardCategory",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// New property: RationDistributionHistory
			Tag:      "rationDistributionHistory",
			Label:    "Ration Distribution History",
			DataType: "rationDistributionHistory",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
	},
}
