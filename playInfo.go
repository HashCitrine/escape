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
	upCoords,downCoords,rightCoords,leftCoords := getAroundCoords(currentCoords)

	playInfo = PlayInfo{
		goalCoords:    goalCoords,
		currentCoords: currentCoords,

		upPlace: getPlaceByCoords(upCoords),
		downPlace: getPlaceByCoords(downCoords),
		rightPlace: getPlaceByCoords(rightCoords),
		leftPlace: getPlaceByCoords(leftCoords),
		inventory : []Code{codeHand},
	}
	return
}

func setPlayInfo(coords Coords/* , item Code */) {
	upCoords,downCoords,rightCoords,leftCoords := getAroundCoords(coords)

	playInfo.upPlace = getPlaceByCoords(upCoords)
	playInfo.downPlace = getPlaceByCoords(downCoords)
	playInfo.rightPlace = getPlaceByCoords(rightCoords)
	playInfo.leftPlace = getPlaceByCoords(leftCoords)

	/* if(item.isItem()) {
		inventory := &playInfo.inventory
		*inventory = append(*inventory, item)
	} */
}