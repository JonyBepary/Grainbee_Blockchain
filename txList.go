package main

import (
	txdefs "github.com/hyperledger-labs/cc-tools-demo/chaincode/txdefs"

	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,
	tx.DeleteAsset,

	txdefs.IssueRationCard,
	txdefs.UpdateMemberInfo,
	txdefs.ReplenishInventory,
	txdefs.CreateRation,
	txdefs.UpdateRation,
	txdefs.ReadRation,
	txdefs.ReadTotalRationsByDistributionPoint,
	txdefs.SetPickupSchedule, // Add the new transaction
	txdefs.GetPickupSchedule, // Add the new transaction
	txdefs.CreateDistributor,
	txdefs.CreateDistributionPoint,
	txdefs.CreateInventory,
	txdefs.SetPickupSchedule, // Add the new transaction
	txdefs.GetPickupSchedule, // Add the new transaction
}

/*
	tx.CreateRation,
	tx.UpdateRation,
	tx.DeleteRation,
	tx.ReadRation,
	tx.ReadRationHistory,

	tx.Search,

	txdefs.CreateRation,
	txdefs.IssueRationCard,
	txdefs.UpdateMemberInfo,
	txdefs.ReplenishInventory,
	txdefs.CreateDistributor,
	txdefs.CreateDistributionPoint,
	txdefs.CreateInventory,
	txdefs.UpdateRation,
	txdefs.DeleteRation,
	txdefs.ReadRation,
	txdefs.ReadRationHistory,
	txdefs.SearchRations,
*/
