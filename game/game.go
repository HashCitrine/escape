package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fieldArray [12][12]Block
var gameInfo GameInfo

// var endGame = false

func init() {
	initAttributeMap()
	initField()
	// initActMap()
	initMovementCommandMap()
	initInteractionCommandMap()
	initPlayInfo(Coords{y: 11, x: 11}, Coords{y: 2, x: 0})
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
		fieldArray[floor.y][floor.x].passable = true
	}

	setAttributeToField()
}

func initPlayInfo(currentCoords Coords, goalCoords Coords) {
	gameInfo = GameInfo{
		goalCoords:   goalCoords,
		playerCoords: currentCoords,
		inventory:    []Component{item.getComponent(codeHand)},
	}

	return
}

func Play() {
	// initGame()
	// clearConsole.print()
	for {
		DrawMap()
		PrintScript()

		scan := Scan()
		if isEnd() {
			break
		}
		clearConsole.print()
		Move(scan)
		Action(scan)
	}
}

func Scan() (scan string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func DrawMap() {
	player := gameInfo.playerCoords

	for i := 0; i < len(fieldArray); i++ {
		var tempArr = fieldArray[i]

		if i > 0 {
			fmt.Println("")
		}
		for j := 0; j < len(tempArr); j++ {
			block := tempArr[j]
			// fmt.Print(block)

			if j > 0 {
				// fmt.Print("\t")
			}

			if i == player.y && j == player.x {
				fmt.Print(playerIcon)
				continue
			}

			if block.isDoor() {
				fmt.Print(doorIcon)
				continue
			}

			/* if len(block.parts) == 0 {
				switch block.passable {
				case true:
					fmt.Print(floorIcon)
				case false:
					fmt.Print(blankIcon)
				}
				continue
			} */

			if !block.passable {
				fmt.Print(blankIcon)
				continue
			}

			if len(block.parts) > 0 {
				fmt.Print(somethingIcon)
				continue
			}

			fmt.Print(floorIcon)

			// fmt.Print(attributeMap[tempArr[j]].icon)
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


	/* if endGame {
		endScript.print()
		return
	} */

	aroundDoor, doorSideDirection := gameInfo.getAroundDoorCoords()
	if aroundDoor != nil {
		doorName := (*aroundDoor).getDoorName()
		lookAtTheDoorscript.print(doorSideDirection, doorName)
		fmt.Println()
	}

	printInventory()
	questionScript.print()
}

func Action(scan string) {
	commandDoor, commandItem := GetDoorAndItem(scan)
	interactionArray := GetInteraction(scan)
	var interactionCode Interaction

	if len(interactionArray) > 0 {
		interactionCode = interactionArray[0]
	}
	place := getPlaceByCoords(gameInfo.playerCoords)

	switch interactionCode {
	case codeAttack :
		// place := getPlaceByCoords(gameInfo.playerCoords)
		/* combatResult := place.combat(scan)
	
		if combatResult {
			return
		} */

		place.combat(scan)
		return
	case codeGet :
		if commandItem.isEmpty() {
			// todo : 무슨 아이템을 주울까요? - script

			// fieldItems := place.findItem()
			// todo : item이 바닥에 있다. - script
			return 
		}	

		place.pickUp(commandItem)
		return
	case codeOpen :
		if len(place.parts) <= 0 {
			break
		}
		part := place.parts[0]
		box := box.getComponent(0)

		if part == box {
			box.Drop()
			return
		}
	}

	// enemy
	/* commandEnemy := GetEnemy(scan)
	if !commandEnemy.isEmpty() {
		if len(interactionArray) == 0 {
			// todo : script - 000이 앞에 있다.
		}

	} */

	// door
	ifDoor, _ := gameInfo.getAroundDoorCoords()
	ifDoorIsOpen := ifDoor != nil && (*ifDoor).isOpen()

	if commandDoor.code == codeWoodDoor && commandItem.isEmpty() {
		commandItem = item.getComponent(codeHand)
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
			alreadyOpenDoorScript.print((*ifDoor).getDoorName())
			return
		}

		aroundDoor, _ := gameInfo.getAroundDoorCoords()

		if aroundDoor == nil {
			canNotFindAroundDoor.print(commandDoor.getName(), commandItem.getName())
			return
		}

		canOpenDoor := commandDoor.tryOpenDoor(commandItem)
		if (*aroundDoor).getDoor() == commandDoor && canOpenDoor {
			// door check
			if len(interactionArray) > 0 {
				for _, tempInteractionCode := range interactionArray {
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

	directionCoords := newCoords(gameInfo.playerCoords, moveCoords)
	if !directionCoords.isPassable() {
		blankScript.print(directionName)
		return
	}

	directionPlace := getPlaceByCoords(directionCoords)

	if (*directionPlace).isOpen() {
		if directionCoords == gameInfo.goalCoords {
			// endGame = true
			updatePlayerPlace(directionCoords /* , directionPlace */)
			return
		}

		doorName := (*directionPlace).getDoorName()
		passDoorScript.print(doorName)

		directionCoords = newCoords(directionCoords, moveCoords)
		directionPlace = getPlaceByCoords(directionCoords)

		updatePlayerPlace(directionCoords /* , directionPlace */)
		return
	}

	if (*directionPlace).isDoor() {
		doorName := (*directionPlace).getDoorName()
		closeDoorScript.print(doorName)
		return
	}

	moveScript.print(directionName)

	items := directionPlace.findItem()
	if len(items) > 0 {
		for _, item := range items {
			itemName := item.getName()
			findItmeScript.print(itemName)
		}

		// todo : 아이템 줍기 별도 구현 필요
		// getItemScript.print(itemName)
		// inventory := &playInfo.inventory
		// *inventory = append(*inventory, (*directionPlace))
		// (*directionPlace) = Component{passable: true}
	}

	// 현재 위치 업데이트
	updatePlayerPlace(directionCoords /* , directionPlace */)
}


func isEnd() bool {
	return gameInfo.goalCoords == gameInfo.playerCoords
}