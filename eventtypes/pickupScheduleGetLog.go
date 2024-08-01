package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var PickupScheduleGetLog = events.Event{
	Tag:         "pickupScheduleGetLog",
	Label:       "Pickup Schedule Get Log",
	Description: "Log of getting the pickup schedule",
	Type:        events.EventLog,
	BaseLog:     "Pickup schedule retrieved",
	Receivers:   []string{"$org1MSP", "$orgMSP"},
}
