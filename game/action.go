////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

type Action interface {
	JSONAction() JSONAction
}

type StealAction struct {
	Direction Point
}

type MeleeAttackAction struct {
	Direction Point
}

type CollectAction struct {
	Direction Point
}

type MoveAction struct {
	Direction Point
}

type UpgradeAction struct {
	Upgrade Upgrade
}

type PurchaseAction struct {
	Item Item
}

type HealAction struct{}

func CreateStealAction(x int, y int) StealAction {
	return StealAction{
		Direction: Point{X: x, Y: y},
	}
}

func CreateMeleeAttackAction(x int, y int) MeleeAttackAction {
	return MeleeAttackAction{
		Direction: Point{X: x, Y: y},
	}
}

func CreateCollectAction(x int, y int) CollectAction {
	return CollectAction{
		Direction: Point{X: x, Y: y},
	}
}

func CreateMoveAction(x int, y int) MoveAction {
	return MoveAction{
		Direction: Point{X: x, Y: y},
	}
}

type JSONAction struct {
	Name    string `json:"ActionName"`
	Content string `json:"Content"`
}

func (a StealAction) JSONAction() JSONAction {
	return JSONAction{
		Name:    "StealAction",
		Content: a.Direction.ToJSON(),
	}
}

func (a MeleeAttackAction) JSONAction() JSONAction {
	return JSONAction{
		Name:    "MeleeAttackAction",
		Content: a.Direction.ToJSON(),
	}
}

func (a CollectAction) JSONAction() JSONAction {
	return JSONAction{
		Name:    "CollectAction",
		Content: a.Direction.ToJSON(),
	}
}

func (a MoveAction) JSONAction() JSONAction {
	return JSONAction{
		Name:    "MoveAction",
		Content: a.Direction.ToJSON(),
	}
}

func (a UpgradeAction) JSONAction() JSONAction {
	return JSONAction{
		Name:    "UpgradeAction",
		Content: "TODO",
	}
}

func (a PurchaseAction) JSONAction() JSONAction {
	return JSONAction{
		Name:    "PurchaseAction",
		Content: "TODO",
	}
}

func (a HealAction) JSONAction() JSONAction {
	return JSONAction{
		Name:    "HealAction",
		Content: "TODO",
	}
}
