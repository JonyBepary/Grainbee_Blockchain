package main

import (
	"github.com/hyperledger-labs/cc-tools-demo/chaincode/assettypes"
	"github.com/hyperledger-labs/cc-tools/assets"
)

var assetTypeList = []assets.AssetType{
	assettypes.Member,
	assettypes.DistributorAsset,
	assettypes.DistributionPoint,
	assettypes.Inventory,
	assettypes.RationAsset,
	assettypes.Secret,
}
