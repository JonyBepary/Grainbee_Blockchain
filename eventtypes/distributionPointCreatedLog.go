package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var DistributionPointCreatedLog = events.Event{
	Tag:         "distributionPointCreatedLog",
	Label:       "Distribution Point Created Log",
	Description: "Log of distribution point creation",
	Type:        events.EventLog,
	BaseLog:     "New distribution point created",
	Receivers:   []string{"$org1MSP", "$orgMSP"},
}
