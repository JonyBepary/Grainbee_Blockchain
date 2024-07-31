package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

// Description of a Inventory as a collection of books
var Inventory = assets.AssetType{
	Tag:         "inventory",
	Label:       "Inventory",
	Description: "Inventory as a collection of books",

	Props: []assets.AssetProp{
		{
			// Primary Key
			Required: true,
			IsKey:    true,
			Tag:      "name",
			Label:    "Inventory Name",
			DataType: "string",
			Writers:  []string{`org3MSP`, "orgMSP"}, // This means only org3 can create the asset (others can edit)
		},
		{
			// Asset reference list
			Tag:      "rations",
			Label:    "Ration Collection",
			DataType: "[]->ration",
		},
		{
			// Asset reference list
			Tag:      "entranceCode",
			Label:    "Entrance Code for the Inventory",
			DataType: "->secret",
		},
	},
}
