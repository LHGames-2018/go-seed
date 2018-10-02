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

func (u Upgrade) ToJSON() string {
	var mapping = make(map[Upgrade]string)
	mapping[UpgradeCarryingCapacity] = "CarryingCapacity"
	mapping[UpgradeMaximumHealth] = "Maximumhealth"
	mapping[UpgradeCollectingSpeed] = "CollectingSpeed"
	mapping[UpgradeAttack] = "AttackPower"
	mapping[UpgradeDefence] = "Defence"
	return mapping[u]
}

type Upgrades struct {
	CarryingCapacity uint
	MaximumHealth    uint
	CollectingSpeed  uint
	Attack           uint
	Defence          uint
}
