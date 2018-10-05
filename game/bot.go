package main

import (
	"fmt"
)

type Bot struct {
	/* add custom struct fields */
}

func (a *Bot) ExecuteAction(player *Player, gamemap *Map) Action {
	fmt.Println(player.Upgrades.Attack)
	gamemap.Print()

	house := gamemap.GetTile(9, 9)
	if house != nil {
		fmt.Println(house.GetPosition().ToJSON())
	} else {
		fmt.Println("failed")
	}

	for _, resource := range gamemap.Resources {
		fmt.Println(resource.Position.ToJSON())
	}

	return CreateMoveAction(0, 1)
}
