package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var InventoryCreatedLog = events.Event{
	Tag:         "inventoryCreatedLog",
	Label:       "Inventory Created Log",
	Description: "Log of inventory creation",
	Type:        events.EventLog,
	BaseLog:     "New inventory created",
	Receivers:   []string{"$org3MSP", "$orgMSP"},
}
