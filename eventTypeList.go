package main

import (
	"github.com/hyperledger-labs/cc-tools-demo/chaincode/eventtypes"
	"github.com/hyperledger-labs/cc-tools/events"
)

var eventTypeList = []events.Event{
	eventtypes.CreateLibraryLog,
	eventtypes.RationCardIssuedLog, // Add the new RationCardIssuedLog event
	eventtypes.MemberInfoUpdatedLog,
	eventtypes.RationCreatedLog,
	eventtypes.RationUpdatedLog,
	eventtypes.DistributorCreatedLog,
	eventtypes.DistributionPointCreatedLog,
	eventtypes.InventoryCreatedLog,
	eventtypes.RationUpdatedLog,
	eventtypes.PickupScheduleSetLog, // Add the new event
	eventtypes.PickupScheduleGetLog, // Add the new event
	eventtypes.RationDeletedLog,
	eventtypes.RationPurchasedLog,
}
