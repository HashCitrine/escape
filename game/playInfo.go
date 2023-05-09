package game

import "fmt"

type PlayInfo struct {
	goalCoords    Coords
	currentCoords Coords

	upPlace    *Code
	downPlace  *Code
	rightPlace *Code
	leftPlace  *Code

	inventory []Code
}

func setPlayInfo(coords Coords) {
	upCoords, downCoords, rightCoords, leftCoords := GetAroundCoords(coords)

	playInfo.upPlace = getPlaceByCoords(upCoords)
	playInfo.downPlace = getPlaceByCoords(downCoords)
	playInfo.rightPlace = getPlaceByCoords(rightCoords)
	playInfo.leftPlace = getPlaceByCoords(leftCoords)
}

func updatePlayerPlace(tempPlaceCoords Coords, tempPlace *Code) {
	currentPlace := getPlaceByCoords(playInfo.currentCoords)
	playInfo.currentCoords = tempPlaceCoords

	*currentPlace = codeFloor
	*tempPlace = codePlayer

	setPlayInfo(playInfo.currentCoords)
}

func (platInfo PlayInfo) getAroundDoorCoords() (*Code, string) {
	aroundPlace := []*Code{playInfo.upPlace, playInfo.downPlace, playInfo.rightPlace, playInfo.leftPlace}
	for i, place := range aroundPlace {
		if place != nil && (*place).isDoor() {
			return place, getDoorSideWayByIndex(i)
		}
	}

	return nil, ""
}

func getDoorSideWayByIndex(index int) string {
	switch index {
	case 0:
		return actMap[upAct].name
	case 1:
		return actMap[downAct].name
	case 2:
		return actMap[rightAct].name
	case 3:
		return actMap[leftAct].name
	}

	return ""
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
