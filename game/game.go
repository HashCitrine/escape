package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var fieldArray [12][13]Block
var player Player

// var endGame = false

func init() {
	initAttributeMap()
	initField()
	// initActMap()
	initMovementCommandMap()
	initInteractionCommandMap()
	InitDropItemMap()
	initEnenyMap()
	initEquipmentMap()
	initPlayer(Coords{y: 11, x: 12}, Coords{y: 2, x: 0})
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

func Play() {
	// initGame()
	// clearConsole.print()
	for {
		DrawMap()

		if isEnd() {
			break
		} else {
			PrintScript()
		}

		scan := Scan()
		clearConsole.print()
		if Action(scan) {
			continue
		}
		Move(scan)
	}
}

func DrawMap() {
	player := player.currentCoords

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

	aroundDoor, doorSideDirection := player.getAroundDoorCoords()
	if aroundDoor != nil {
		doorName := (*aroundDoor).getDoorName()
		lookAtTheDoorscript.print(doorSideDirection, doorName)
		fmt.Println()
	}

	printInventory()
	questionScript.print()
	playerInfoScript.print(player.hp)
}

func Action(scan string) bool {
	commandDoor, commandItem := GetDoorAndItem(scan)
	interactionArray := GetInteraction(scan)
	var interactionCode Interaction

	if len(interactionArray) > 0 {
		interactionCode = interactionArray[0]
	}
	place := getPlaceByCoords(player.currentCoords)

	fmt.Println("commandDoor, commandItem, interactionCode : ", commandDoor, commandItem, interactionCode)

	switch interactionCode {
	case codeWear:
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
	ifDoor, _ := player.getAroundDoorCoords()
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

		aroundDoor, _ := player.getAroundDoorCoords()

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

	directionCoords := newCoords(player.currentCoords, moveCoords)
	if !directionCoords.isPassable() {
		blankScript.print(directionName)
		return
	}

	directionPlace := getPlaceByCoords(directionCoords)

	if (*directionPlace).isOpen() {
		if directionCoords == player.goalCoords {
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

	if player.hp <= 0 {
		deadScript.print()
		return true
	}

	if player.goalCoords == player.currentCoords {
		endScript.print()
		return true
	}

	return false
}

func (block *Block) combat(code Interaction) bool {
	blockEnemy := enemyMap[block]

	if blockEnemy.component.isEmpty() {
		// todo : 적이 없다. - script
		return false
	}

	enemyOffence := blockEnemy.common.offence
	defense := player.getDefense()

	switch code {
	case codeRun:
		rand.Seed(time.Now().UnixNano())
		result := rand.Float64() * 100

		if result > 0.5 {
			return false
		}
	case codeAttack:
		playerOffence := player.common.offence
		rightAttack := player.rightHand.getOffense()
		leftAttack := player.leftHand.getOffense()

		playerOffence += rightAttack + leftAttack
		blockEnemy.hp -= playerOffence
		enemyMap[block] = blockEnemy

		attackedScript.print(blockEnemy.component.getName(), playerOffence)
		enemyInfoScript.print(blockEnemy.component.getName(), blockEnemy.hp)

		if blockEnemy.hp <= 0 {
			// todo : enenmy가 죽었다. - script
			blockEnemy.component.Drop()
			parts := (*block).parts
			for i, part := range parts {
				if part == blockEnemy.component {
					(*block).parts = append(parts[:i], parts[i+1:]...)
				}
			}
			return true
		}
	case codeShield:
		defense += player.rightHand.getDefense() + player.leftHand.getDefense()
	case codeRecovery:
		if useItem(item.getComponent(codePortion)) {
			blockEnemy.hp += 30
		}
	}

	enemyOffence -= defense

	if enemyOffence < 0 {
		enemyOffence = 0
	}

	player.hp -= enemyOffence
	attackedScript.print("고양이", enemyOffence)
	return true
}
