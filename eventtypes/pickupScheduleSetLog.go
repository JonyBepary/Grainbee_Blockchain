package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var PickupScheduleSetLog = events.Event{
	Tag:         "pickupScheduleSetLog",
	Label:       "Pickup Schedule Set Log",
	Description: "Log of setting the pickup schedule",
	Type:        events.EventLog,
	BaseLog:     "Pickup schedule set",
	Receivers:   []string{"$org1MSP", "$orgMSP"},
}
