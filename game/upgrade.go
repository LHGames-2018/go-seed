////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

const (
	UpgradeCarryingCapacity = 0
	UpgradeMaximumHealth    = 3
	UpgradeCollectingSpeed  = 4
	UpgradeAttack           = 1
	UpgradeDefence          = 2
)

type Upgrade int

type Upgrades struct {
	CarryingCapacity uint
	MaximumHealth    uint
	CollectingSpeed  uint
	Attack           uint
	Defence          uint
}
