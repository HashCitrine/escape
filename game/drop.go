package game

import (
	"math/rand"
	"time"
)

type DropItem struct {
	item        Component
	amount      int
	probability float64
}

func (component Component) Drop() {
	dropTable := dropItemMap[component]
	place := getPlaceByCoords(player.currentCoords)

	rand.Seed(time.Now().UnixNano())
	result := rand.Float64() * 100

	for _, drop := range dropTable {
		result -= drop.probability

		if result <= 0 {
			dropItem := drop.item
			componentArray := (*place).parts
			(*place).parts = append(componentArray, dropItem)

			if dropItem.code == 0 {
				noDropScript.print()
				return
			}

			dropItemScript.print(dropItem.getName())
			break
		}
	}

	return
}

func (block *Block) pickUp(item Component) {
	parts := (*block).parts

	for i, component := range parts {
		if component == item {
			player.inventory = append(player.inventory, item)
			(*block).parts = append(parts[:i], parts[i+1:]...)

			// todo : 00 아이템을 주웠다. - script
			return
		}
	}
}

func opening(interaction Interaction, door Component, item Component) bool {
	if interaction.isCanDo(door, item) {
		return useItem(item)
	}

	return false
}

func useItem(item Component) bool {
	inventory := player.inventory
	for i, haveItem := range inventory {
		if item == haveItem {
			player.inventory = append(inventory[:i], inventory[i+1:]...)
			return true
		}
	}
	return false
}
