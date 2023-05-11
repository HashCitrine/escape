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
	// initActMap()
	initMovingCommandMap()
	initPlayInfo(Coords{y: 4, x: 1}, Coords{y: 0, x: 7})
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

	// setPlayInfo(currentCoords)
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
		scan := scanner.Text()

		if endGame {
			break
		}
		clearConsole.print()
		Move(scan)
		Action(scan)
		// Action(scan)
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

func Action( /* door Code, item Code, */ scan string) {
	var door Code
	var item Code
	var act Code

	for code, attribute := range attributeMap {
		if code.isOpenDoor() {
			continue
		}

		for _, command := range attribute.commands {
			if strings.Contains(scan, command) {
				if code.isDoor() {
					door = code
				}

				if code.isItem() {
					item = code
				}

				if code.isActioning() {
					act = code
				}
			}
		}
	}

	ifDoor, _ := playInfo.getAroundDoorCoords()
	ifDoorIsOpen := ifDoor != nil && (*ifDoor).isOpenDoor()

	/* if door > 0 && !checkActToDoor(acts, door, ifDoor, ifDoorIsOpen) {
		return
	} */

	if item > 0 && !checkInventory(item) {
		notHaveItemScript.print(item.getName())
		return
	}

	if door == codeWoodDoor && item == 0 {
		item = codeHand
	}

	if door > 0 && item > 0 {
		if ifDoor != nil && ifDoorIsOpen {
			alreadyOpenDoorScript.print((*ifDoor).getName())
			return
		}

		openingDoor := door + item
		aroundDoor, _ := playInfo.getAroundDoorCoords()

		if *aroundDoor == door && openingDoor.isOpenDoor() {
			if act.isCanActioning(door, item) {
				useItemToDoorScript.print(item.getName(), door.getName())
				*aroundDoor = openingDoor
				return
			}

			doNotActToDoorScript.print(door.getName())
			return
		}

		canNotUseItemToDoorScript.print(item.getName(), door.getName())
	}
}

func Move(act string) {
	var moving Moving

	for command, code := range movingCommandMap {
		if strings.Contains(act, string(command)) {
			moving = code
		}
	}

	if moving == 0 {
		return
	}

	moveCoords := moving.getDirectionInfo()
	direction := moving.getDirectionName()

	directionCoords := newCoords(playInfo.currentCoords, moveCoords)
	if checkOutFieldByCoords(directionCoords) || *getPlaceByCoords(directionCoords) == codeBlank {
		blankScript.print()
		return
	}

	directionPlace := getPlaceByCoords(directionCoords)

	if (*directionPlace).isOpenDoor() {
		if directionCoords == playInfo.goalCoords {
			endGame = true
			updatePlayerPlace(directionCoords, directionPlace)
			return
		}
		// doorName := attributeMap[Code((*directionPlace).getDoorNumber())*door].getName()
		doorName := (*directionPlace).getName()
		passDoorScript.print(doorName)

		directionCoords = newCoords(directionCoords, moveCoords)
		directionPlace = getPlaceByCoords(directionCoords)

		updatePlayerPlace(directionCoords, directionPlace)
		return
	}

	if (*directionPlace).isDoor() {
		doorName := attributeMap[(*directionPlace)].getName()
		closeDoorScript.print(doorName)
		return
	}

	moveScript.print(direction)

	if (*directionPlace).isItem() {
		itemName := attributeMap[(*directionPlace)].getName()
		findItmeScript.print(itemName)
		inventory := &playInfo.inventory
		*inventory = append(*inventory, (*directionPlace))
	}

	// 현재 위치 업데이트
	updatePlayerPlace(directionCoords, directionPlace)
}
