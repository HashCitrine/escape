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

func Action(scan string) {
	var door Code
	var item Code
	var interact Code

	for code, attribute := range attributeMap {
		if code.isOpen() {
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
					interact = code
				}
			}
		}
	}

	ifDoor, _ := playInfo.getAroundDoorCoords()
	ifDoorIsOpen := ifDoor != nil && (*ifDoor).isOpen()

	if door == codeWoodDoor && item == 0 {
		item = codeHand
	}

	if item > 0 {
		ishaveItem := checkInventory(item)

		switch ishaveItem {
		case true:
			if door == 0 {
				doNotActByItemScript.print(item.getName())
			}
		case false:
			notHaveItemScript.print(item.getName())
		}
	}

	if door > 0 && item == 0 {
		needActAnythingScript.print()
	}

	if door > 0 && item > 0 {
		if ifDoor != nil && ifDoorIsOpen {
			alreadyOpenDoorScript.print((*ifDoor).getName())
			return
		}

		tryOpenTheDoor := door + item
		aroundDoor, _ := playInfo.getAroundDoorCoords()

		if aroundDoor == nil {
			canNotFindAroundDoor.print(door.getName(), item.getName())
			return
		}

		if *aroundDoor == door && tryOpenTheDoor.isOpen() {
			if interact.isCanActioning(door, item) {
				useItemToDoorScript.print(item.getName(), door.getName())
				*aroundDoor = tryOpenTheDoor
				return
			}

			doNotActToDoorScript.print(door.getName())
			return
		}

		canNotUseItemToDoorScript.print(item.getName(), door.getName())
	}
}

func Move(scan string) {
	var movement Movement

	for command, code := range movingCommandMap {
		if strings.Contains(scan, string(command)) {
			movement = code
		}
	}

	if movement == 0 {
		return
	}

	moveCoords := movement.getDirectionInfo()
	direction := movement.getDirectionName()

	directionCoords := newCoords(playInfo.currentCoords, moveCoords)
	if checkOutFieldByCoords(directionCoords) || *getPlaceByCoords(directionCoords) == codeBlank {
		blankScript.print()
		return
	}

	directionPlace := getPlaceByCoords(directionCoords)

	if (*directionPlace).isOpen() {
		if directionCoords == playInfo.goalCoords {
			endGame = true
			updatePlayerPlace(directionCoords, directionPlace)
			return
		}
		
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
