package game

import "fmt"

type Act string
type ActCommand string

type Acting struct {
	name      string
	direction Code
	coords    Coords
}

const (
	upAct        Act = "up"
	downAct      Act = "down"
	rightAct     Act = "right"
	leftAct      Act = "left"
	openAct      Act = "open"
	breakOpenAct Act = "breakOpen"
	keyOpenAct   Act = "keyOpen"
)

func (act Act) getActing() Acting {
	return actMap[act]
}

func (move Acting) actMove() {
	// var move = act[0]
	if move.direction == codeFloor {
		directionCoords := NewCoords(playInfo.currentCoords, move.coords)
		if checkOutFieldByCoords(directionCoords) || *getPlaceByCoords(directionCoords) == codeBlank {
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

			directionCoords = NewCoords(directionCoords, move.coords)
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

func actByAttribute(door Code, item Code, acts []Act) {
	ifDoor, _ := playInfo.getAroundDoorCoords()
	ifDoorIsOpen := ifDoor != nil && (*ifDoor).isOpenDoor()
	if door > 0 && !checkActToDoor(acts, door, ifDoor, ifDoorIsOpen) {
		return
	}

	if item > 0 && !checkInventory(item) {
		fmt.Printf("%s를 가지고 있지 않다. 다른 방법을 찾아보자.\n", attributeMap[item].getName())
		return
	}

	if door == codeWoodDoor && item == 0 {
		item = codeHand
	}

	if door > 0 && item > 0 {
		if ifDoor != nil && ifDoorIsOpen {
			fmt.Printf("%s은 이미 열려있다. 지나갈 수 있을 것 같다.\n", attributeMap[(*ifDoor)].getName())
			return
		}

		openingDoor := door + item
		aroundDoor, _ := playInfo.getAroundDoorCoords()

		if *aroundDoor == door && openingDoor.isOpenDoor() {
			fmt.Printf("%s(으)로 %s을 열었다. 이제 지나갈 수 있다.\n", attributeMap[item].getName(), attributeMap[door].getName())
			*aroundDoor = openingDoor
		} else {
			fmt.Printf("%s(으)로는 %s을 열 수 없다. 다른 방법을 찾아보자.\n", attributeMap[item].getName(), attributeMap[door].getName())
		}
	}
}
