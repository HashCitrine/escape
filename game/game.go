package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fieldArray [12][13]Block
var gameInfo GameInfo

// var endGame = false

func init() {
	initAttributeMap()
	initField()
	// initActMap()
	initMovementCommandMap()
	initInteractionCommandMap()
	InitDropItemMap()
	initEnenyMap()
	initPlayInfo(Coords{y: 11, x: 12}, Coords{y: 2, x: 0})
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
		/* player: Player{
			charactor: Charactor{
				hp: 50,
				common: TacticalCommon{
					offence: 3,
					defense: 0,
				},
			},
		}, */
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
		if Action(scan) {
			continue
		}
		Move(scan)
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
	playerInfoScript.print(player.charactor.hp)
}

func Action(scan string) bool {
	commandDoor, commandItem := GetDoorAndItem(scan)
	interactionArray := GetInteraction(scan)
	var interactionCode Interaction

	if len(interactionArray) > 0 {
		interactionCode = interactionArray[0]
	}
	place := getPlaceByCoords(gameInfo.playerCoords)

	switch interactionCode {
	case codeWear :
		commandItem.wear()
		return true
	case codeAttack, codeRun, codeShield, codeRecovery:
		// place := getPlaceByCoords(gameInfo.playerCoords)
		/* combatResult := place.combat(scan)

		if combatResult {
			return
		} */

		place.combat(interactionCode)
		return true
	case codeGet:
		if commandItem.isEmpty() {
			// todo : 무슨 아이템을 주울까요? - script

			// fieldItems := place.findItem()
			// todo : item이 바닥에 있다. - script
			return true
		}

		place.pickUp(commandItem)
		return true
	case codeOpen:
		if len(place.parts) <= 0 {
			break
		}

		part := place.parts[0]
		closeBox := box.getComponent(codeCloseBox)
		openBox := box.getComponent(codeOpenBox)

		if part == closeBox {
			closeBox.Drop()
			place.parts[0] = openBox
			return true
		}

		if part == openBox {
			openBoxScript.print()
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
			return true
		}

		aroundDoor, _ := gameInfo.getAroundDoorCoords()

		if aroundDoor == nil {
			canNotFindAroundDoor.print(commandDoor.getName(), commandItem.getName())
			return true
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
						return true
					}
				}
			} else {
				return true
			}

			/* if interaction.isCanDo(commandDoor, commandItem) {
				useItemToDoorScript.print(commandItem.getName(), commandDoor.getName())
				*aroundDoor = tryOpenTheDoor
				return
			} */

			doNotActToDoorScript.print(commandDoor.getName())
			return true
		}

		canNotUseItemToDoorScript.print(commandItem.getName(), commandDoor.getName())
	}

	return false
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
			directionPlace.printParts()
			return
		}

		doorName := (*directionPlace).getDoorName()
		passDoorScript.print(doorName)

		directionCoords = newCoords(directionCoords, moveCoords)
		directionPlace = getPlaceByCoords(directionCoords)

		updatePlayerPlace(directionCoords /* , directionPlace */)
		directionPlace.printParts()
		return
	}

	if (*directionPlace).isDoor() {
		doorName := (*directionPlace).getDoorName()
		closeDoorScript.print(doorName)
		return
	}

	moveScript.print(directionName)

	/* items := directionPlace.findItem()
	if len(items) > 0 {
		for _, item := range items {
			itemName := item.getName()
			findItmeScript.print(itemName)
		}
	} */

	// 현재 위치 업데이트
	updatePlayerPlace(directionCoords /* , directionPlace */)
	directionPlace.printParts()
}

func isEnd() bool {

	if player.charactor.hp <= 0 {
		deadScript.print()
		return true
	}

	if gameInfo.goalCoords == gameInfo.playerCoords {
		endScript.print()
		return true
	}

	return false
}
