package game

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
	initGame()
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
	for i := 0; i < len(fieldArray); i++ {
		var tempArr = fieldArray[i]

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

	aroundDoor := playInfo.getAroundDoorCoords()
	if aroundDoor != nil {
		doorName := attributeMap[(*aroundDoor)].getName()
		fmt.Printf("%s이 보인다.", doorName)
		fmt.Println()
		fmt.Println()
	}

	fmt.Print("소지품 : ")
	var inventoryList string
	for _, item := range playInfo.inventory {
		if len(inventoryList) > 0 {
			inventoryList += ", "
		}

		if item == codeHand {
			continue
		}

		inventoryList += attributeMap[item].getName()
	}

	if len(inventoryList) == 0 {
		inventoryList += "없음"
	}

	fmt.Println(inventoryList)

	fmt.Println()
	fmt.Println()

	fmt.Println("(어떤 행동을 하시겠습니까?)")
}

func Action(chose string) {
	// var act Code
	var door Code
	var item Code
	var acts []Act = nil

	for k, actArray := range actCommandMap {
		if strings.Contains(chose, string(k)) {
			if acts != nil {
				fmt.Println("욕심 부리지 말자. 차근차근 하나씩 행동해야 한다.")
				return
			}

			acts = actArray
		}
	}

	if acts != nil {
		acts[0].getActing().actMove()
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
