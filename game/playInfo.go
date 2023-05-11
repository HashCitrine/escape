package game

import "fmt"

type PlayInfo struct {
	goalCoords    Coords
	currentCoords Coords

	inventory []Code
}

func updatePlayerPlace(tempPlaceCoords Coords, tempPlace *Code) {
	currentPlace := getPlaceByCoords(playInfo.currentCoords)
	playInfo.currentCoords = tempPlaceCoords

	*currentPlace = codeFloor
	*tempPlace = codePlayer
}

func (playInfo PlayInfo) getAroundDoorCoords() (*Code, string) {
	coordsArray := getAroundCoords(playInfo.currentCoords)

	for i, coords := range coordsArray {
		place := getPlaceByCoords(coords)
		if place != nil && (*place).isDoor() {
			return place, Movement((i + 1) * moving).getDirectionName()
		}
	}

	return nil, ""
}

func printInventory() {
	inventorys := playInfo.inventory

	if len(inventorys) < 2 {
		return
	}

	fmt.Print("소지품 : ")

	for i, inventory := range inventorys {
		if i > 1 {
			fmt.Print(", ")
		}

		if inventory == codeHand {
			continue
		}

		fmt.Print(attributeMap[inventory].getName())
	}

	fmt.Println()
	fmt.Println()
}
