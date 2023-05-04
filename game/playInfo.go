package game

import . "escape/game/util"

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
	upCoords,downCoords,rightCoords,leftCoords := GetAroundCoords(coords)

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

func (platInfo PlayInfo) getAroundDoorCoords() *Code {
	aroundPlace := []*Code{playInfo.upPlace, playInfo.downPlace, playInfo.rightPlace, playInfo.leftPlace}
	for _, place := range aroundPlace {
		if place != nil && (*place).isDoor() {
			return place
		}
	}

	return nil
}