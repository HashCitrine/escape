package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var endGame = false

func consoleClear() {
	fmt.Print("\033[H\033[2J")
}

func PlayGame() {
	consoleClear()
	for {
		DrawMap()
		Script()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		chose := scanner.Text()

		if endGame {
			break
		}
		consoleClear()

		Action(chose)
	}
}

func DrawMap() {
	for i := 0; i < len(fieldArr); i++ {
		var tempArr = fieldArr[i]

		if i > 0 {
			fmt.Println("")
		}
		for j := 0; j < len(tempArr); j++ {
			if j > 0 {
				fmt.Print("\t")
			}

			fmt.Print(attributeMap[tempArr[j]].mark)
		}
	}
}

func Script() {
	fmt.Println()
	fmt.Println()
	currentPlace := getPlaceByCoords(playInfo.currentCoords)
	if currentPlace == attributeMap[codePlayer].place[0] {
		fmt.Println("(당신은 미로의 함정에 빠졌습니다. 이곳을 빠져나가야 합니다.)")
		fmt.Println()
	}

	if endGame {
		fmt.Println("(밝은 빛이 보입니다. 당신은 탈출에 성공했습니다.)")
		return
	}

	fmt.Print("소지품 : ")
	for i, inventory := range playInfo.inventory {
		if i > 0 {
			fmt.Print(", ")
		}

		fmt.Print(attributeMap[inventory].getName())
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("(어떤 행동을 하시겠습니까?)")
}

func Action(chose string) {
	// var act Code
	var door Code
	var item Code
	var act []Acting = nil

	for k := range actMap {
		if strings.Contains(chose, string(k)) {
			if act != nil {
				fmt.Println("욕심 부리지 말자. 차근차근 하나씩 행동해야 한다.")
				return
			}

			act = actMap[k]
		}
	}

	if act != nil {
		justMove(act[0])
	}

	// 특정 요소를 지칭하고 명령한 경우
	for code, attribute := range attributeMap {
		if code.isOpenDoor() {
			continue
		}

		for i := range attribute.commands {
			var command = attribute.commands[i]

			if strings.Contains(chose, command) {
				/* switch {
				case code/100 > 0:
					item = Code(code)
				case code/10 > 0:
					door = Code(code)
				} */

				if code.isDoor() {
					door = code
				}

				if code.isItem() {
					item = code
				}
			}
		}
	}

	ifDoor := getAroundDoorCoords()
	ifDoorIsOpen := ifDoor != nil && (*ifDoor).isOpenDoor()
	if door > 0 && !actCheck(act, door, ifDoor, ifDoorIsOpen) {
		return
	}

	if door > 0 && item > 0 {
		if ifDoor != nil && ifDoorIsOpen {
			fmt.Printf("%s은 이미 열려있다. 지나갈 수 있을 것 같다.\n", attributeMap[(*ifDoor)].getName())
			return
		}

		if !checkInventory(item) {
			fmt.Printf("%s를 가지고 있지 않다. 다른 방법을 찾아보자.\n", attributeMap[item].getName())
			return
		}

		openingDoor := door + item
		aroundDoor := getAroundDoorCoords()

		if *aroundDoor == door && openingDoor.isOpenDoor() {
			fmt.Printf("%s(으)로 %s을 열었다. 이제 지나갈 수 있다.\n", attributeMap[item].getName(), attributeMap[door].getName())
			*aroundDoor = openingDoor
		} else {
			fmt.Printf("%s(으)로는 %s을 열 수 없다. 다른 방법을 찾아보자.\n", attributeMap[item].getName(), attributeMap[door].getName())
		}
	}
}

func justMove(move Acting) {
	// var move = act[0]
	if move.direction == codeFloor {
		directionCoords := newCoords(playInfo.currentCoords, move.coords)
		if checkOutField(directionCoords) || *getPlaceByCoords(directionCoords) == codeBlank {
			// directionPlace = *tempPlace
			fmt.Println("막힌 길이다. 다시 생각해보자.")
			return
		}

		directionPlace := getPlaceByCoords(directionCoords)

		if (*directionPlace).isOpenDoor() {
			if directionCoords == playInfo.goalCoords {
				endGame = true
				updatePlayerPlace(directionCoords, directionPlace)
				return
			}
			doorName := attributeMap[Code((*directionPlace).getDoorNumber())*door].getName()
			fmt.Printf("%s을 지나왔다.\n", doorName)

			directionCoords = newCoords(directionCoords, move.coords)
			directionPlace = getPlaceByCoords(directionCoords)

			updatePlayerPlace(directionCoords, directionPlace)
			return
		}

		if (*directionPlace).isDoor() {
			doorName := attributeMap[(*directionPlace)].getName()
			fmt.Printf("%s이 닫혀 있다. 이대로는 나아갈 수 없다.\n", doorName)
			return
		}

		fmt.Printf("%s로 이동했다.\n", move.name)
		
		if (*directionPlace).isItem() {
			itemName := attributeMap[(*directionPlace)].getName()
			fmt.Printf("%s가 떨어져 있다. 어딘가에 사용할 수 있을 것 같다. 챙겨놓도록 하자.\n", itemName)
			inventory := &playInfo.inventory
			*inventory = append(*inventory, (*directionPlace))
		}

		// 현재 위치 업데이트
		updatePlayerPlace(directionCoords, directionPlace)
	}
}

func updatePlayerPlace(tempPlaceCoords Coords, tempPlace *Code) {
	currentPlace := getPlaceByCoords(playInfo.currentCoords)
	playInfo.currentCoords = tempPlaceCoords

	*currentPlace = codeFloor
	*tempPlace = codePlayer

	setPlayInfo(playInfo.currentCoords)
}

func checkInventory(item Code) bool {
	for _, haveItem := range playInfo.inventory {
		if item == haveItem {
			return true
		}
	}
	return false
}

func actCheck(act []Acting, doorCode Code, ifDoor *Code, ifDoorIsOpen bool) bool {
	if act != nil {
		for _, acting := range act {
			if acting.direction == doorCode {
				return true
			}
		}
	}

	if ifDoorIsOpen {
		fmt.Printf("%s은 이미 열려있다. 지나갈 수 있을 것 같다.\n", attributeMap[(*ifDoor)].getName())
	}

	fmt.Printf("%s은 굳게 닫혀있다.\n", attributeMap[(*ifDoor)].getName())

	return false
}
