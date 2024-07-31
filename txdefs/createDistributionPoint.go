package txdefs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger-labs/cc-tools-demo/chaincode/datatypes"
	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/events"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var CreateDistributionPoint = tx.Transaction{
	Tag:         "createDistributionPoint",
	Label:       "Create Distribution Point",
	Description: "Create a new distribution point",
	Method:      "POST",
	Callers: []accesscontrol.Caller{ // Only org1 admin can call this transaction
		{
			MSP: "org1MSP",
			OU:  "admin",
		},
		{
			MSP: "orgMSP",
			OU:  "admin",
		},
	},
	Args: []tx.Argument{
		{
			Tag:         "distributionPointId",
			Label:       "Distribution Point ID",
			Description: "Distribution Point ID",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "name",
			Label:       "Name of the distribution point",
			Description: "Name of the distribution point",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "address",
			Label:       "Address",
			Description: "Address",
			DataType:    "address",
			Required:    false,
		},
		{
			Tag:         "coordinates",
			Label:       "GPS Location or Coordinates",
			Description: "GPS Location or Coordinates",
			DataType:    "coordinates",
			Required:    false,
		},
		{
			Tag:         "contactInformation",
			Label:       "Contact Information",
			Description: "Contact Information",
			DataType:    "contactInfo",
			Required:    false,
		},
		{
			Tag:         "distributor",
			Label:       "Distributor Object",
			Description: "Distributor Object",
			DataType:    "->distributor",
			Required:    false,
		},
		{
			Tag:         "operatingHours",
			Label:       "Operating Hours",
			Description: "Operating Hours",
			DataType:    "operatingHours",
			Required:    false,
		},
		{
			Tag:         "capacity",
			Label:       "Capacity",
			Description: "Capacity",
			DataType:    "integer",
			Required:    false,
		},

		{
			Tag:         "lastInspectionDate",
			Label:       "Last Inspection Date",
			Description: "Last Inspection Date",
			DataType:    "datetime",
			Required:    false,
		},
		{
			Tag:         "inspectionStatus",
			Label:       "Inspection Status",
			Description: "Inspection Status",
			DataType:    "inspectionStatus",
			Required:    false,
		},
		{
			Tag:         "numberOfCounters",
			Label:       "Number of Counters",
			Description: "Number of Counters",
			DataType:    "integer",
			Required:    false,
		},
		{
			Tag:         "inventory",
			Label:       "Inventory",
			Description: "Inventory",
			DataType:    "inventory",
			Required:    false,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		distributionPointId, _ := req["distributionPointId"].(string)
		name, _ := req["name"].(string)
		address, _ := req["address"].(datatypes.Address)             // This is a custom data type
		coordinates, _ := req["coordinates"].(datatypes.Coordinates) // This is a custom data type
		contactInformation, _ := req["contactInformation"].(datatypes.ContactInfo)

		operatingHours, _ := req["operatingHours"].(datatypes.OperatingHours) // This is a custom data type
		capacity, _ := req["capacity"].(int)
		lastInspectionDate, _ := req["lastInspectionDate"].(time.Time)
		inspectionStatus, _ := req["inspectionStatus"].(datatypes.InspectionStatus)
		numberOfCounters, _ := req["numberOfCounters"].(int)
		inventory, _ := req["inventory"].(string)

		distributionPointMap := make(map[string]interface{})
		distributionPointMap["@assetType"] = "distributionPoint"
		distributionPointMap["distributionPointId"] = distributionPointId
		distributionPointMap["name"] = name
		distributionPointMap["address"] = address
		distributionPointMap["coordinates"] = coordinates
		distributionPointMap["contactInformation"] = contactInformation
		// reference to the distributor object
		distributorKey, _ := req["distributor"].(assets.Key)

		if distributorKey != nil {
			// Retrieve the distributor asset
			distributorAsset, err := distributorKey.Get(stub)
			if err != nil {
				return nil, errors.WrapError(err, "failed to get distributor asset from the ledger")
			}
			distributionPointMap["distributor"] = distributorAsset
		}

		distributionPointMap["operatingHours"] = operatingHours
		distributionPointMap["capacity"] = capacity
		distributionPointMap["lastInspectionDate"] = lastInspectionDate
		distributionPointMap["inspectionStatus"] = inspectionStatus
		distributionPointMap["numberOfCounters"] = numberOfCounters
		distributionPointMap["inventory"] = inventory

		distributionPointAsset, err := assets.NewAsset(distributionPointMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = distributionPointAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		distributionPointJSON, nerr := json.Marshal(distributionPointAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		logMsg, ok := json.Marshal(fmt.Sprintf("New distribution point created: %s", distributionPointId))
		if ok != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		events.CallEvent(stub, "distributionPointCreatedLog", logMsg)

		return distributionPointJSON, nil
	},
}
