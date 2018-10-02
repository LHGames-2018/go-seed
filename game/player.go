////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

type Player struct {
	HealthCurrent    uint
	HealthMax        uint
	ResourceCurrent  uint
	ResourceCapacity uint
	ResourceTotal    uint
	Attack           uint
	Defence          uint
	Score            uint
	CollectingSpeed  float32
	Name             string
	Position         Point
	House            Point
	Upgrades         Upgrades
	Items            []Item
}

type JSONPlayer struct {
	HealthCurrent    uint      `json:"Health"`
	HealthMax        uint      `json:"MaxHealth"`
	ResourceCurrent  uint      `json:"CarriedResources"`
	ResourceCapacity uint      `json:"CarryingCapacity"`
	ResourceTotal    uint      `json:"TotalResource"`
	Attack           uint      `json:"AttackPower"`
	Defence          uint      `json:"Defence"`
	Score            uint      `json:"Score"`
	CollectingSpeed  float32   `json:"CollectingSpeed"`
	Name             string    `json:"Name"`
	Position         JSONPoint `json:"Position"`
	House            JSONPoint `json:"HouseLocation"`
	Upgrades         []uint    `json:"UpgradeLevels"`
	Items            []uint    `json:"CarriedItems"`
}

func (json JSONPlayer) Player() Player {
	return Player{
		HealthCurrent:    json.HealthCurrent,
		HealthMax:        json.HealthMax,
		ResourceCurrent:  json.ResourceCurrent,
		ResourceCapacity: json.ResourceCapacity,
		ResourceTotal:    json.ResourceTotal,
		Attack:           json.Attack,
		Defence:          json.Defence,
		Score:            json.Score,
		CollectingSpeed:  json.CollectingSpeed,
		Name:             json.Name,
		Position:         json.Position.Point(),
		House:            json.House.Point(),
		Upgrades: Upgrades{
			CarryingCapacity: json.Upgrades[UpgradeCarryingCapacity],
			MaximumHealth:    json.Upgrades[UpgradeMaximumHealth],
			CollectingSpeed:  json.Upgrades[UpgradeCollectingSpeed],
			Attack:           json.Upgrades[UpgradeAttack],
			Defence:          json.Upgrades[UpgradeDefence],
		},
		Items: Items(json.Items).ToItems(),
	}
}
