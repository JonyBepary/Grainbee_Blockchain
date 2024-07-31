package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var RationCardIssuedLog = events.Event{
	Tag:         "rationCardIssuedLog",
	Label:       "Ration Card Issued Log",
	Description: "Log of a ration card issuance",
	Type:        events.EventLog,
	BaseLog:     "New ration card issued",
	Receivers:   []string{"$org1MSP", "$orgMSP"},
}
