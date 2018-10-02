package main

import (
	"fmt"
)

type Bot struct {
	/* add custom struct fields */
}

func (a *Bot) ExecuteAction(Player *Player, Map *Map) Action {
	fmt.Println(Player.Upgrades.Attack)
	Map.Print()

	return CreateMeleeAttackAction(1, 0)
}
