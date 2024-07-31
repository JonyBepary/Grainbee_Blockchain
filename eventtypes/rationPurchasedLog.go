package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var RationPurchasedLog = events.Event{
	Tag:         "rationPurchasedLog",
	Label:       "Ration Purchased Log",
	Description: "Log of ration purchase",
	Type:        events.EventLog,
	BaseLog:     "Ration purchased",
	Receivers:   []string{"$org1MSP", "$orgMSP"},
}
