package game

import . "escape/game/util"

func getPlace(y int, x int) *Code {
	if y < 0 || x < 0 {
		return nil
	}
	return &fieldArray[y][x]
}

func checkOutFieldByCoords(coords Coords) bool {
	if coords.Y < 0 || coords.X < 0 ||
		len(fieldArray) <= coords.Y || len(fieldArray[0]) <= coords.X {
		return true
	}
	return false
}

func getPlaceByCoords(coords Coords) *Code {
	if checkOutFieldByCoords(coords) {
		return nil
	}

	return &fieldArray[coords.Y][coords.X]
}
