package game

import (
	"math/rand"
	"time"
)

type DropItem struct {
	item        Component
	amount      int
	probability int
}

var dropItemMap map[Component][]DropItem

func InitDropItemMap() {
	portionItem := item.getComponent(codePortion)
	woodSwordItem := item.getComponent(codeWoodSword)
	ironSwordItem := item.getComponent(codeIronSword)
	// woodShield := item.getComponent(codeWoodShield)
	leatherRobeItem := item.getComponent(codeLeatherRobe)
	leatherPantsItem := item.getComponent(codeLeatherPants)
	leatherHatItem := item.getComponent(codeLeatherHat)

	portion1 := DropItem{item: portionItem, amount: 1, probability: 15}
	portion2 := DropItem{item: portionItem, amount: 2, probability: 10}
	portion3 := DropItem{item: portionItem, amount: 3, probability: 5}
	woodSword := DropItem{item: woodSwordItem, amount: 1, probability: 20}
	ironSword := DropItem{item: ironSwordItem, amount: 1, probability: 15}
	leatherRobe := DropItem{item: leatherRobeItem, amount: 1, probability: 7}
	leatherPants := DropItem{item: leatherPantsItem, amount: 1, probability: 8}
	leatherHat := DropItem{item: leatherHatItem, amount: 1, probability: 10}
	notDrop := DropItem{probability: 10}
	boxDrop := []DropItem{
		woodSword, ironSword,
		leatherRobe, leatherPants, leatherHat,
		portion1, portion2, portion3,
	}

	portion1 = DropItem{item: portionItem, amount: 1, probability: 70}
	notDrop = DropItem{probability: 30}
	squirrelDrop := []DropItem{
		portion1, notDrop,
	}

	portion1 = DropItem{item: portionItem, amount: 1, probability: 50}
	portion2 = DropItem{item: portionItem, amount: 2, probability: 30}
	notDrop = DropItem{probability: 20}
	rabbitDrop := []DropItem{
		portion1, portion2, notDrop,
	}

	key := DropItem{item: item.getComponent(codeKey), amount: 1, probability: 100}
	deerDrop := []DropItem{
		key,
	}

	dropItemMap = map[Component][]DropItem{
		box.getComponent(0):                    boxDrop,
		enemy.getComponent(Code(codeSquirrel)): squirrelDrop,
		enemy.getComponent(Code(codeRabbit)):   rabbitDrop,
		enemy.getComponent(Code(codeDeer)):     deerDrop,
	}
}

func (component Component) Drop() Component {
	dropTable := dropItemMap[component]

	rand.Seed(time.Now().UnixNano())
	result := rand.Float64() * 100

	for _, drop := range dropTable {
		result -= float64(drop.probability)

		if result <= 0 {
			return drop.item
		}
	}

	return Component{}
}

func (block Block) pickUp(item Component) {
	parts := block.parts

	for _, component := range parts {
		if component == item {
			gameInfo.inventory = append(gameInfo.inventory, item)
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
	inventory := gameInfo.inventory
	for i, haveItem := range inventory {
		if item == haveItem {
			gameInfo.inventory = append(inventory[:i], inventory[i+1:]...)
			return true
		}
	}
	return false
}
