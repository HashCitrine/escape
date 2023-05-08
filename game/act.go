package game

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
	if move.direction == codeFloor {
		directionCoords := NewCoords(playInfo.currentCoords, move.coords)
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

			directionCoords = NewCoords(directionCoords, move.coords)
			directionPlace = getPlaceByCoords(directionCoords)

			updatePlayerPlace(directionCoords, directionPlace)
			return
		}

		if (*directionPlace).isDoor() {
			doorName := attributeMap[(*directionPlace)].getName()
			closeDoorScript.print(doorName)
			return
		}

		moveScript.print(move.name)

		if (*directionPlace).isItem() {
			itemName := attributeMap[(*directionPlace)].getName()
			findItmeScript.print(itemName)
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
		notHaveItemScript.print(item.getName())
		return
	}

	if door == codeWoodDoor && item == 0 {
		item = codeHand
	}

	if door > 0 && item > 0 {
		if ifDoor != nil && ifDoorIsOpen {
			alreadyOpenDoor.print((*ifDoor).getName())
			return
		}

		openingDoor := door + item
		aroundDoor, _ := playInfo.getAroundDoorCoords()

		if *aroundDoor == door && openingDoor.isOpenDoor() {
			useItemToDoor.print(item.getName(), door.getName())
			*aroundDoor = openingDoor
		} else {
			doNotUseItemToDoor.print(item.getName(), door.getName())
		}
	}
}
