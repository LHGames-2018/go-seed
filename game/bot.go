package main

import (
	"fmt"
)

type Bot struct {
	/* add custom struct fields */
}

func (a *Bot) ExecuteAction(Player *Player, Map *Map) {
	fmt.Println(Player.Upgrades.Attack)
	Map.Print()
}
