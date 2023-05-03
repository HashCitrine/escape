package main

import (
	"fmt"
)

// const door = 10
// const item = 100

// var fieldArr = [][]Code{
// 	{codeBlank, codeBlank, codeBlank, codeBlank, codeFloor, codeFloor, codeFloor, codeCloseDoor},
// 	{codeBlank, codeBlank, codeBlank, codeBlank, codeWoodDoor, codeBlank, codeBlank, codeBlank},
// 	{codeHammer, codeFloor, codeGlassDoor, codeFloor, codeFloor, codeBlank, codeBlank, codeBlank},
// 	{codeBlank, codeFloor, codeBlank, codeBlank, codeFloor, codeBlank, codeBlank, codeBlank},
// 	{codeBlank, codeStart, codeBlank, codeBlank, codeFloor, codeFloor, codeBlank, codeBlank},
// 	{codeBlank, codeBlank, codeBlank, codeBlank, codeBlank, codeKey, codeBlank, codeBlank},
// }

// type Item struct {
// 	*Common
// }

// type Door struct {
// 	*Common
// }

// type Structor struct {
// 	*Common
// }

// var start Structor
// var floor Structor
// var blank Structor

// var closeDoor Door
// var glassDoor Door
// var woodDoor Door

// var key Item
// var hammer Item
// var hand Item

// var currentPlaceCoords = Coords{y: 4, x: 1}
// var endPlaceCoords = Coords{y: 0, x: 7}

func initGame() {

	initField()
	initAttributeMap()
	initActMap()
	initPlayInfo(Coords{y: 4, x: 1}, Coords{y: 0, x: 7})
	/* for i := range key.commands {
		key := Act(key.commands[i])
		actMap[key] = keyOpen
	} */

	// setAttributeToField()
}

func getMark(mark interface{}) string {
	return "[" + fmt.Sprintf("%v", mark) + "]"
}

func getStringArray(stringArray ...string) []string {
	return stringArray
}

/* func GetCodePointerArray(codeAddressArray ...*Code) []*Code {
	return codeAddressArray
} */

/* func GetCodePointerArray2(codeAddressArray ...[]int) []*Code {
	result := make([]*Code, len(codeAddressArray))

	for i := range result {
		y := codeAddressArray[i][0]
		x := codeAddressArray[i][1]
		result[i] = &fieldArr[y][x]
	}

	return result
} */
