package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fieldArray [6][8]Component
var playInfo GameInfo

var endGame = false

func init() {
	initAttributeMap()
	initField()
	// initActMap()
	initMovementCommandMap()
	initInteractionCommandMap()
	initPlayInfo(Coords{y: 4, x: 1}, Coords{y: 0, x: 7})
}

func initField() {
	/* fieldArray = make([][]Component, 6)

	for i := range fieldArray {
		fieldArray[i] = make([]Component, 8)
		for j := range fieldArray[i] {
			fieldArray[i][j] = codeBlank
		}
	} */

	for _, floor := range floorPlace {
		fieldArray[floor.y][floor.x] = Component{passable: true}
	}

	setAttributeToField()
}

func initPlayInfo(currentCoords Coords, goalCoords Coords) {
	playInfo = GameInfo{
		goalCoords:   goalCoords,
		playerCoords: currentCoords,
		inventory:    []Component{item.getComponent(codeHand, true)},
	}

	return
}

func Play() {
	// initGame()
	// clearConsole.print()
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
	player := playInfo.playerCoords

	for i := 0; i < len(fieldArray); i++ {
		var tempArr = fieldArray[i]

		if i > 0 {
			fmt.Println("")
		}
		for j := 0; j < len(tempArr); j++ {
			if j > 0 {
				// fmt.Print("\t")
			}

			if i == player.y && j == player.x {
				fmt.Print(playerIcon)
				continue
			}

			component := tempArr[j]
			if component.passable == false && component.codetype != door {
				fmt.Print(blankIcon)
				continue
			}

			if component.codetype == "" && component.code == 0 {
				fmt.Print(floorIcon)
				continue
			}

			fmt.Print(attributeMap[tempArr[j]].icon)
		}
	}
}

func PrintScript() {
	fmt.Println()
	fmt.Println()
	/* currentPlace := getPlaceByCoords(playInfo.playerCoords)
	if currentPlace == attributeMap[codePlayer].place[0] {
		startScript.print()
	} */

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
	var commandDoor Component
	var commandItem Component

	for component, attribute := range attributeMap {
		if component.isOpen() {
			continue
		}

		for _, command := range attribute.commands {
			if strings.Contains(scan, command) {
				if component.isDoor() {
					commandDoor = component
				}

				if component.isItem() {
					commandItem = component
				}
			}
		}
	}

	ifDoor, _ := playInfo.getAroundDoorCoords()
	ifDoorIsOpen := ifDoor != nil && (*ifDoor).isOpen()

	if commandDoor.code == codeWoodDoor && commandItem.isEmpty() {
		commandItem = item.getComponent(codeHand, true)
	}

	if !commandItem.isEmpty() {
		switch hasItem(commandItem) {
		case true:
			if commandDoor.isEmpty() {
				doNotActByItemScript.print(commandItem.getName())
			}
		case false:
			notHaveItemScript.print(commandItem.getName())
		}
	}

	if !commandDoor.isEmpty() && commandItem.isEmpty() {
		needActAnythingScript.print()
	}

	if !commandDoor.isEmpty() && !commandItem.isEmpty() {
		if ifDoor != nil && ifDoorIsOpen {
			alreadyOpenDoorScript.print((*ifDoor).getName())
			return
		}

		aroundDoor, _ := playInfo.getAroundDoorCoords()

		if aroundDoor == nil {
			canNotFindAroundDoor.print(commandDoor.getName(), commandItem.getName())
			return
		}

		canOpenDoor := commandDoor.tryOpenDoor(commandItem)
		if *aroundDoor == commandDoor && canOpenDoor {
			var interactionArray []Interaction
			// var interaction Interaction
			for command, interactions := range interactionCommandMap {
				if strings.Contains(scan, string(command)) {
					interactionArray = interactions
				}
			}

			// door check
			if len(interactionArray) > 0 {
				for _, tempInteractionCode := range interactionArray {
					fmt.Println(tempInteractionCode.isCanDo(commandDoor, commandItem))
					if tempInteractionCode.isCanDo(commandDoor, commandItem) {
						// interaction = tempInteractionCode
						useItemToDoorScript.print(commandItem.getName(), commandDoor.getName())
						(*aroundDoor).passable = true
						return
					}
				}
			} else {
				return
			}

			/* if interaction.isCanDo(commandDoor, commandItem) {
				useItemToDoorScript.print(commandItem.getName(), commandDoor.getName())
				*aroundDoor = tryOpenTheDoor
				return
			} */

			doNotActToDoorScript.print(commandDoor.getName())
			return
		}

		canNotUseItemToDoorScript.print(commandItem.getName(), commandDoor.getName())
	}
}

func Move(scan string) {
	var movement Movement

	for command, code := range movementCommandMap {
		if strings.Contains(scan, string(command)) {
			movement = code
		}
	}

	if movement == 0 {
		return
	}

	moveCoords, directionName := movement.getDirectionInfo()
	// directionName := movement.getDirectionName()

	directionCoords := newCoords(playInfo.playerCoords, moveCoords)
	if !directionCoords.isPassable() {
		blankScript.print(directionName)
		return
	}

	directionPlace := getPlaceByCoords(directionCoords)

	if (*directionPlace).isOpen() {
		if directionCoords == playInfo.goalCoords {
			endGame = true
			updatePlayerPlace(directionCoords /* , directionPlace */)
			return
		}

		doorName := (*directionPlace).getName()
		passDoorScript.print(doorName)

		directionCoords = newCoords(directionCoords, moveCoords)
		directionPlace = getPlaceByCoords(directionCoords)

		updatePlayerPlace(directionCoords /* , directionPlace */)
		return
	}

	if (*directionPlace).isDoor() {
		doorName := attributeMap[(*directionPlace)].getName()
		closeDoorScript.print(doorName)
		return
	}

	moveScript.print(directionName)

	if (*directionPlace).isItem() {
		itemName := attributeMap[(*directionPlace)].getName()
		findItmeScript.print(itemName)
		getItemScript.print(itemName)
		inventory := &playInfo.inventory
		*inventory = append(*inventory, (*directionPlace))
		(*directionPlace) = Component{passable: true}
	}

	// 현재 위치 업데이트
	updatePlayerPlace(directionCoords /* , directionPlace */)
}
