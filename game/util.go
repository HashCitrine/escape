package game

import (
	"bufio"
	"os"
)

func getPlace(y int, x int) *Block {
	if y < 0 || x < 0 {
		return nil
	}
	return &fieldArray[y][x]
}

func getPlaceByCoords(coords Coords) *Block {
	if checkOutFieldByCoords(coords) {
		return nil
	}

	return &fieldArray[coords.y][coords.x]
}

func checkOutFieldByCoords(coords Coords) bool {
	if coords.y < 0 || coords.x < 0 ||
		len(fieldArray) <= coords.y || len(fieldArray[0]) <= coords.x {
		return true
	}
	return false
}

func Scan() (scan string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}