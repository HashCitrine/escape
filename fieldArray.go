package main

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
	if coords.y < 0 || coords.x < 0 {
		return nil
	}
	return &fieldArr[coords.y][coords.x]
}

func checkInField(coords Coords) bool {
	if coords.y < 0 || coords.x < 0 ||
		len(fieldArr) <= coords.y || len(fieldArr[0]) <= coords.x {
		return false
	}
	return true
}
