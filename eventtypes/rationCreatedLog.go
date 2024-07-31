package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var RationCreatedLog = events.Event{
	Tag:         "rationCreatedLog",
	Label:       "Ration Created Log",
	Description: "Log of ration creation",
	Type:        events.EventLog,
	BaseLog:     "New ration created",
	Receivers:   []string{"$org2MSP", "$orgMSP"},
}
