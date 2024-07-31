package eventtypes

import (
	"github.com/hyperledger-labs/cc-tools/events"
)

var MemberInfoUpdatedLog = events.Event{
	Tag:         "memberInfoUpdatedLog",
	Label:       "Member Info Updated Log",
	Description: "Log of member information update",
	Type:        events.EventLog,
	BaseLog:     "Member information updated",
	Receivers:   []string{"$org1MSP", "$orgMSP"},
}
