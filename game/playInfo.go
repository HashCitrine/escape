package game

import "fmt"

type PlayInfo struct {
	goalCoords    Coords
	currentCoords Coords

/* 	upPlace    *Code
	downPlace  *Code
	rightPlace *Code
	leftPlace  *Code */

	inventory []Code
}

/* func setPlayInfo(coords Coords) {
	upCoords, downCoords, rightCoords, leftCoords := getAroundCoords(coords)

	playInfo.upPlace = getPlaceByCoords(upCoords)
	playInfo.downPlace = getPlaceByCoords(downCoords)
	playInfo.rightPlace = getPlaceByCoords(rightCoords)
	playInfo.leftPlace = getPlaceByCoords(leftCoords)
} */

func updatePlayerPlace(tempPlaceCoords Coords, tempPlace *Code) {
	currentPlace := getPlaceByCoords(playInfo.currentCoords)
	playInfo.currentCoords = tempPlaceCoords

	*currentPlace = codeFloor
	*tempPlace = codePlayer

	// setPlayInfo(playInfo.currentCoords)
}

func (platInfo PlayInfo) getAroundDoorCoords() (*Code, string) {
	// aroundPlace := []*Code{playInfo.upPlace, playInfo.downPlace, playInfo.rightPlace, playInfo.leftPlace}
	coordsArray := getAroundCoords(platInfo.currentCoords)

	for i, coords := range coordsArray {
		place := getPlaceByCoords(coords)
		if place != nil && (*place).isDoor() {
			return place, Moving((i + 1) * moving).getDirectionName()
		}
	}

	return nil, ""
}

/* func getDoorSideWayByIndex(index int) Act {
	switch index {
	case 0:
		return upAct
	case 1:
		return downAct
	case 2:
		return rightAct
	case 3:
		return leftAct
	}

	return ""
} */

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
