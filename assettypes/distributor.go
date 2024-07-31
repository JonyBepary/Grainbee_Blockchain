package assettypes

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/hyperledger-labs/cc-tools-demo/chaincode/datatypes"
	"github.com/hyperledger-labs/cc-tools/assets"
)

type Distributor struct {
	DistributorID      string                `json:"distributorId"`
	Name               string                `json:"name"`
	Address            datatypes.Address     `json:"address"`
	ContactInformation datatypes.ContactInfo `json:"contactInformation"`
	LicenseNumber      string                `json:"licenseNumber"`
	LicenseIssueDate   string                `json:"licenseIssueDate"`
	LicenseExpiryDate  string                `json:"licenseExpiryDate"`
	DistributionArea   string                `json:"distributionArea"`
	LastInspectionDate string                `json:"lastInspectionDate"`
}

var DistributorAsset = assets.AssetType{
	Tag:         "distributor",
	Label:       "Distributor",
	Description: "Information about a ration distributor",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "distributorId",
			Label:    "Distributor ID",
			DataType: "string",                      // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Name of the distributor",
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
		},
		{
			// Optional property
			Tag:      "contactInformation",
			Label:    "Contact Information",
			DataType: "contactInfo",
		},
		{
			// Optional property
			Tag:      "licenseNumber",
			Label:    "License Number",
			DataType: "string",
			Validate: func(licenseNumber interface{}) error {
				licenseNumberStr, ok := licenseNumber.(string)
				if !ok {
					return fmt.Errorf("license number must be a string")
				}

				errChan := make(chan error, 8) // Buffer size matches the number of checks
				var wg sync.WaitGroup

				// Basic checks
				wg.Add(3)
				go func() {
					defer wg.Done()
					if licenseNumberStr == "" {
						errChan <- fmt.Errorf("license number must be non-empty")
					}
				}()

				go func() {
					defer wg.Done()
					if len(licenseNumberStr) != 14 {
						errChan <- fmt.Errorf("license number must be 14 characters long")
					}
				}()

				go func() {
					defer wg.Done()
					licenseParts := strings.Split(licenseNumberStr, "-")
					if len(licenseParts) != 2 {
						errChan <- fmt.Errorf("license number must contain a hyphen")
					}
				}()

				// Advanced checks
				licenseParts := strings.Split(licenseNumberStr, "-")
				if len(licenseParts) == 2 {
					wg.Add(5)
					go func() {
						defer wg.Done()
						if len(licenseParts[0]) != 4 {
							errChan <- fmt.Errorf("license number first part must be 4 characters long")
						}
					}()

					go func() {
						defer wg.Done()
						if licenseParts[0] != "DCLN" {
							errChan <- fmt.Errorf("license number first part must be DCLN")
						}
					}()

					go func() {
						defer wg.Done()
						if licenseParts[1] == "" {
							errChan <- fmt.Errorf("license number second part must be non-empty")
						}
					}()

					go func() {
						defer wg.Done()
						if _, err := strconv.Atoi(licenseParts[1]); err != nil {
							errChan <- fmt.Errorf("license number second part must be a number")
						}
					}()

					go func() {
						defer wg.Done()
						if len(licenseParts[1]) != 9 {
							errChan <- fmt.Errorf("license number second part must be 9 characters long")
						}
					}()
				}

				// Wait for all goroutines to finish
				go func() {
					wg.Wait()
					close(errChan)
				}()

				// Collect errors
				var errors []string
				for err := range errChan {
					errors = append(errors, err.Error())

				}

				if len(errors) > 0 {
					return fmt.Errorf("validation failed: %s", strings.Join(errors, "; "))
				}

				return nil
			},
		},
		{
			// Optional property
			Tag:      "licenseIssueDate",
			Label:    "License Issue Date",
			DataType: "datetime",
		},
		{
			// Optional property
			Tag:      "licenseExpiryDate",
			Label:    "License Expiry Date",
			DataType: "datetime",
		},
		{
			// Optional property
			Tag:      "distributionArea",
			Label:    "Distribution Area",
			DataType: "string",
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Optional property
			Tag:      "lastInspectionDate",
			Label:    "Last Inspection Date",
			DataType: "datetime",
		},
	},
}
