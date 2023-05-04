package util

type Coords struct {
	Y int
	X int
}

func GetAroundCoords(coords Coords) (upCoords Coords, downCoords Coords, rightCoords Coords, leftCoords Coords) {
	upCoords, downCoords, rightCoords, leftCoords = coords, coords, coords, coords

	upCoords.Y = coords.Y - 1
	downCoords.Y = coords.Y + 1
	leftCoords.X = coords.X - 1
	rightCoords.X = coords.X + 1

	return
}

func NewCoords(currentCoords Coords, addCoords Coords) Coords {
	currentCoords.X += addCoords.X
	currentCoords.Y += addCoords.Y

	return currentCoords
}