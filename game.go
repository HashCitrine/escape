package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PlayGame() {
	for {
		if EndGame() {
			break
		}
		DrawMap()
		Script()

		// var chose string
		// fmt.Scan(&chose)
		// fmt.Println("test : ", chose)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		chose := scanner.Text()
		fmt.Print("\033[H\033[2J") // clear

		Action(chose)
	}
}

func DrawMap() {
	for i := 0; i < len(fieldArr); i++ {
		var tempArr = fieldArr[i]

		if i > 0 {
			fmt.Println("")
		}
		// fmt.Print("\t")
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
	currentPlace := getPlaceByCoords(playInfo.currentCoords)
	if currentPlace == attributeMap[codePlayer].place[0] {
		fmt.Println("(당신은 미로의 함정에 빠졌습니다. 이곳을 빠져나가야 합니다.)")
	}

	if currentPlace == attributeMap[codeCloseDoor].place[0] {
		fmt.Println("(밝은 빛이 보입니다. 당신은 탈출에 성공했습니다.)")
	}

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

	if act == nil {
		fmt.Println("정신 차리자. 무엇이든 행동해야 한다.")
		return
	}

	// 문
	if len(act) > 1 {
		// 1. 문 이름 구하기 attributeMap
		fmt.Println("")
	}

	// 단순 이동
	justMove(act[0])

	// 특정 요소를 지칭하고 명령한 경우
	for code, attribute := range attributeMap {
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

	// 2. 명령한 요소가 실행될 수 있는 환경인지 + act 명령에 적합한지 확인 후 실행
	/* if item > 0 && door > 0 && item/100 == door/10 {

	} */
	if door > 0 && item > 0 {
		openingDoor := door + item
		aroundDoor := getAroundDoorCoords()
		fmt.Println(aroundDoor)
		if *aroundDoor == door && (openingDoor).isOpenDoor() {
			// currentPlace := getPlaceByCoords(playInfo.currentCoords)

			/* switch door {
			case *playInfo.upPlace:
				*playInfo.upPlace = openingDoor
			case *playInfo.downPlace:
				*playInfo.downPlace = openingDoor
			case *playInfo.rightPlace:
				*playInfo.rightPlace = openingDoor
			case *playInfo.leftPlace:
				*playInfo.leftPlace = openingDoor
			} */
			*aroundDoor = openingDoor
		}
	}

	//

}

func justMove(move Acting) {
	// var move = act[0]
	if move.direction == codeFloor {
		tempPlaceCoords := playInfo.currentCoords

		tempPlaceCoords.y = tempPlaceCoords.y + move.coords.y
		tempPlaceCoords.x = tempPlaceCoords.x + move.coords.x

		// 확인 필요 → field 이상의 값의 좌표가 들어왔을 경우
		var directionPlace Code
		var tempPlace *Code

		// tempCheck := checkInField(tempPlaceCoords)
		// fmt.Println(tempCheck)	// debug
		if checkInField(tempPlaceCoords) {
			tempPlace = getPlaceByCoords(tempPlaceCoords)
			directionPlace = *tempPlace
		} else {
			directionPlace = codeBlank
		}

		/* tempPlace := getPlaceByCoords(tempPlaceCoords)
		directionPlace := *tempPlace */
		// 확인 필요

		// debug
		// fmt.Println("test : ", tempPlaceCoords)
		// fmt.Println("directionAttribute : ", directionPlace)
		// debug

		if directionPlace == codeBlank {
			fmt.Println("막힌 길이다. 다시 생각해보자.")
			return
		}

		if directionPlace.isDoor() {
			doorName := attributeMap[directionPlace].commands[0]
			fmt.Println(doorName, "이 닫혀 있다. 이대로는 나아갈 수 없다.")
			return
		}

		if directionPlace.isItem() {
			itemName := attributeMap[directionPlace].commands[0]
			fmt.Println(itemName, "가 떨어져 있다. 어딘가에 사용할 수 있을 것 같다. 챙겨놓도록 하자.")
			// playInfo.currentCoords = tempPlaceCoords
			// *tempPlace = codePlayer
			// return
			inventory := &playInfo.inventory
			*inventory = append(*inventory, directionPlace)
			// fmt.Println(playInfo.inventory) // debug
			// updatePlayerPlace(tempPlaceCoords, tempPlace)
			// return
		}

		// 현재 위치 업데이트
		updatePlayerPlace(tempPlaceCoords, tempPlace)
	}
}

func updatePlayerPlace(tempPlaceCoords Coords, tempPlace *Code) {
	currentPlace := getPlaceByCoords(playInfo.currentCoords)
	*currentPlace = codeFloor

	playInfo.currentCoords = tempPlaceCoords
	*tempPlace = codePlayer

	setPlayInfo(playInfo.currentCoords)
}

func EndGame() bool {
	if playInfo.currentCoords == playInfo.goalCoords {
		fmt.Println()
		fmt.Println("1", playInfo.currentCoords)
		fmt.Print("2", playInfo.goalCoords)
		return true
	}

	return false
}

//

/* func AttributeNilCheck(attribute Attribute) bool {
	return len(attribute.commands) == 0 && attribute.field == 0 && attribute.mark == ""
} */

/* func GetName(attribute int) (name string) {
	switch attribute {
	case -1:
		name = "start"
	case 0:
		name = "floor"
	case 1:
		name = "blank"
	case 10:
		name = "closeDoor"
	case 20:
		name = "woodDoor"
	case 30:
		name = "glassDoor"
	case 100:
		name = "key"
	case 200:
		name = "hammer"
	case 300:
		name = "open"
	}

	return name
} */

/* func GetField(attribute int) (field string) {
	switch attribute {
	case codeStart:
		field = "[1]"
	case codeFloor:
		field = "[ ]"
	case codeBlank:
		field = ""
	case codeCloseDoor:
		field = "|♠|"
	case codeGlassDoor:
		field = "|♣|"
	case codeWoodDoor:
		field = "|♥|"
	case codeKey:
		field = "[§]"
	case codeHammer:
		field = "[↔]"
	case codeHand:
		field = "[※]"
	}

	return field
}
*/
