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

func (i Item) ToJSON() string {
	var mapping = make(map[Item]string)
	mapping[ItemSword] = "Sword"
	mapping[ItemShield] = "Shield"
	mapping[ItemBackpack] = "Backpack"
	mapping[ItemPickaxe] = "Pickaxe"
	mapping[ItemHealthPotion] = "HealthPotion"
	return mapping[i]
}

type Items []uint

func (i Items) ToItems() []Item {
	var items []Item
	for _, item := range i {
		items = append(items, Item(item))
	}
	return items
}
