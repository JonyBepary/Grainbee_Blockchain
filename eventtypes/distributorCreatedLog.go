package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var DistributorCreatedLog = events.Event{
	Tag:         "distributorCreatedLog",
	Label:       "Distributor Created Log",
	Description: "Log of distributor creation",
	Type:        events.EventLog,
	BaseLog:     "New distributor created",
	Receivers:   []string{"$org1MSP", "$orgMSP"},
}
