package game

import "fmt"

type GameInfo struct {
	goalCoords   Coords
	playerName   string
	playerCoords Coords
	inventory    []Component
	// player       Player
}

func updatePlayerPlace(tempPlaceCoords Coords /* , tempPlace *Code */) {
	gameInfo.playerCoords = tempPlaceCoords
}

func (playInfo GameInfo) getAroundDoorCoords() (*Block, string) {
	coordsArray := getAroundCoords(playInfo.playerCoords)

	for i, coords := range coordsArray {
		place := getPlaceByCoords(coords)
		if place != nil && (*place).isDoor() {
			_, directionName := Movement(i + 1).getDirectionInfo()
			return place, directionName
		}
	}

	return nil, ""
}

func hasItem(item Component) bool {
	for _, hasItem := range gameInfo.inventory {
		if hasItem.equals(item) {
			return true
		}
	}
	return false
}

func printInventory() {
	inventorys := gameInfo.inventory

	if len(inventorys) < 2 {
		return
	}

	fmt.Print("소지품 : ")

	hand := item.getComponent(codeHand)
	for i, inventory := range inventorys {
		if i > 1 {
			fmt.Print(", ")
		}

		if inventory == hand {
			continue
		}

		fmt.Print(attributeMap[inventory].getName())
	}

	fmt.Println()
	fmt.Println()
}
