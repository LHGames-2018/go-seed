////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

const (
	ItemSword        = 0
	ItemShield       = 1
	ItemBackpack     = 2
	ItemPickaxe      = 3
	ItemHealthPotion = 4
)

type Item uint

type Items []uint

func (i Items) ToItems() []Item {
	var items []Item
	for _, item := range i {
		items = append(items, Item(item))
	}
	return items
}
