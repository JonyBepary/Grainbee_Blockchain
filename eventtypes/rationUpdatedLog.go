package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var RationUpdatedLog = events.Event{
	Tag:         "rationUpdatedLog",
	Label:       "Ration Updated Log",
	Description: "Log of ration update",
	Type:        events.EventLog,
	BaseLog:     "Ration updated",
	Receivers:   []string{"$org2MSP", "$orgMSP"},
}
