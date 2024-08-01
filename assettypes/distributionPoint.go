package assettypes

import (
	"fmt"

	"github.com/hyperledger-labs/cc-tools/assets"
)

var DistributionPoint = assets.AssetType{
	Tag:         "distributionPoint",
	Label:       "Distribution Point",
	Description: "Information about a ration distribution point",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "distributionPointId",
			Label:    "Distribution Point ID",
			DataType: "string",                      // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Name of the distribution point",
			DataType: "string",
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
			Tag:      "address",
			Label:    "Address",
			DataType: "address",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
		{
			// Optional property
			Tag:      "coordinates",
			Label:    "GPS Location or Coordinates",
			DataType: "coordinates",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
		{
			// Optional property
			Tag:      "contactInformation",
			Label:    "Contact Information",
			DataType: "contactInfo",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
		{
			// Optional property
			Tag:      "distributor",
			Label:    "Distributor Object",
			DataType: "->distributor",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
		{
			// Optional property
			Tag:      "operatingHours",
			Label:    "Operating Hours",
			DataType: "operatingHours",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
		{
			// Optional property
			Tag:      "capacity",
			Label:    "Capacity",
			DataType: "integer",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
		{
			// Optional property
			Tag:      "Distributionstatus",
			Label:    "Status",
			DataType: "string",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
		{
			// Optional property
			Tag:      "lastInspectionDate",
			Label:    "Last Inspection Date",
			DataType: "datetime",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
		{
			// Optional property
			Tag:      "numberOfCounter",
			Label:    "Number of Counter",
			DataType: "integer",

			Writers: []string{`org1MSP`, "orgMSP"},
		},
		{
			// Optional property
			Tag:      "inspectionStatus",
			Label:    "Inspection Status",
			DataType: "inspectionStatus",
		},
		{
			// Optional property
			Tag:      "inventory",
			Label:    "Inventory",
			DataType: "string",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
	},
}
