package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var RationDeletedLog = events.Event{
	Tag:         "rationDeletedLog",
	Label:       "Ration Deleted Log",
	Description: "Log of ration deletion",
	Type:        events.EventLog,
	BaseLog:     "Ration deleted",
	Receivers:   []string{"$org2MSP", "$orgMSP"},
}
