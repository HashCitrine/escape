package main

type Coords struct {
	y int
	x int
}

var fieldArr [][]Code

func initField() {
	fieldArr = make([][]Code, 6)

	for i := range fieldArr {
		fieldArr[i] = make([]Code, 8)
		for j := range fieldArr[i] {
			fieldArr[i][j] = codeBlank
		}
	}
}

func getPlace(y int, x int) *Code {
	if y < 0 || x < 0 {
		return nil
	}
	return &fieldArr[y][x]
}

func getPlaceByCoords(coords Coords) *Code {
	if checkOutField(coords) {
		return nil
	}

	return &fieldArr[coords.y][coords.x]
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

func checkOutField(coords Coords) bool {
	if coords.y < 0 || coords.x < 0 ||
		len(fieldArr) <= coords.y || len(fieldArr[0]) <= coords.x {
		return true
	}
	return false
}
