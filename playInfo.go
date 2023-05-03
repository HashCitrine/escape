package main

type PlayInfo struct {
	goalCoords    Coords
	currentCoords Coords

	upPlace    *Code
	downPlace  *Code
	rightPlace *Code
	leftPlace  *Code

	inventory []Code
}

var playInfo PlayInfo

func initPlayInfo(currentCoords Coords, goalCoords Coords) {
	// upCoords,downCoords,rightCoords,leftCoords := getAroundCoords(currentCoords)

	playInfo = PlayInfo{
		goalCoords:    goalCoords,
		currentCoords: currentCoords,
		inventory : []Code{codeHand},
	}

	setPlayInfo(currentCoords)
	return
}

func setPlayInfo(coords Coords) {
	upCoords,downCoords,rightCoords,leftCoords := getAroundCoords(coords)

	playInfo.upPlace = getPlaceByCoords(upCoords)
	playInfo.downPlace = getPlaceByCoords(downCoords)
	playInfo.rightPlace = getPlaceByCoords(rightCoords)
	playInfo.leftPlace = getPlaceByCoords(leftCoords)
}