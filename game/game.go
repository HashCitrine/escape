package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fieldArray [][]Code
var playInfo PlayInfo

var endGame = false

func init() {
	initField()
	initAttributeMap()
	initActMap()
	initActCommandMap()
	initPlayInfo(Coords{Y: 4, X: 1}, Coords{Y: 0, X: 7})
}

func initField() {
	fieldArray = make([][]Code, 6)

	for i := range fieldArray {
		fieldArray[i] = make([]Code, 8)
		for j := range fieldArray[i] {
			fieldArray[i][j] = codeBlank
		}
	}
}

func initPlayInfo(currentCoords Coords, goalCoords Coords) {
	playInfo = PlayInfo{
		goalCoords:    goalCoords,
		currentCoords: currentCoords,
		inventory:     []Code{codeHand},
	}

	setPlayInfo(currentCoords)
	return
}

func PlayGame() {
	// initGame()
	clearConsole.print()
	for {
		DrawMap()
		PrintScript()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		chose := scanner.Text()

		if endGame {
			break
		}
		clearConsole.print()

		Action(chose)
	}
}

func DrawMap() {
	for i := 0; i < len(fieldArray); i++ {
		var tempArr = fieldArray[i]

		if i > 0 {
			fmt.Println("")
		}
		for j := 0; j < len(tempArr); j++ {
			if j > 0 {
				// fmt.Print("\t")
			}

			fmt.Print(attributeMap[tempArr[j]].icon)
		}
	}
}

func PrintScript() {
	fmt.Println()
	fmt.Println()
	currentPlace := getPlaceByCoords(playInfo.currentCoords)
	if currentPlace == attributeMap[codePlayer].place[0] {
		startScript.print()
	}

	if endGame {
		endScript.print()
		return
	}

	aroundDoor, doorSideDirection := playInfo.getAroundDoorCoords()
	if aroundDoor != nil {
		doorName := (*aroundDoor).getName()
		lookAtTheDoorscript.print(doorSideDirection, doorName)
		fmt.Println()
	}

	printInventory()
	questionScript.print()
}

func Action(chose string) {
	// var act Code
	var door Code
	var item Code
	var acts []Act = nil

	for k, actArray := range actCommandMap {
		if strings.Contains(chose, string(k)) {
			if acts != nil {
				doNotTooManyActScript.print()
				return
			}

			acts = actArray
		}
	}

	if acts != nil {
		move(acts)
	}

	// 특정 요소를 지칭하고 명령한 경우
	for code, attribute := range attributeMap {
		if code.isOpenDoor() {
			continue
		}

		for i := range attribute.commands {
			var command = attribute.commands[i]

			if strings.Contains(chose, command) {
				if code.isDoor() {
					door = code
				}

				if code.isItem() {
					item = code
				}
			}
		}
	}

	actByAttribute(door, item, acts)
}
