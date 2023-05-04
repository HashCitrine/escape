package main

type Coords struct {
	y int
	x int
}

func getAroundCoords(coords Coords) (upCoords Coords, downCoords Coords, rightCoords Coords, leftCoords Coords) {
	upCoords, downCoords, rightCoords, leftCoords = coords, coords, coords, coords

	upCoords.y = coords.y - 1
	downCoords.y = coords.y + 1
	leftCoords.x = coords.x - 1
	rightCoords.x = coords.x + 1

	return
}

func newCoords(currentCoords Coords, addCoords Coords) Coords {
	currentCoords.x += addCoords.x
	currentCoords.y += addCoords.y

	return currentCoords
}