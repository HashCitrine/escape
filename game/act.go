package game

type Act string
type ActCommand string

type Acting struct {
	target Code
	coords Coords
}

var actCommandMap map[ActCommand][]Act
var actMap map[Act]Acting

const (
	upAct        Act = "위"
	downAct      Act = "아래"
	rightAct     Act = "오른쪽"
	leftAct      Act = "왼쪽"
	openAct      Act = "열기"
	breakOpenAct Act = "부숴서열기"
	keyOpenAct   Act = "열쇠열기"
)

func initActMap() {
	upCoords, downCoords, rightCoords, leftCoords := GetAroundCoords(Coords{})
	up := Acting{target: codeFloor, coords: upCoords}
	down := Acting{target: codeFloor, coords: downCoords}
	right := Acting{target: codeFloor, coords: rightCoords}
	left := Acting{target: codeFloor, coords: leftCoords}

	open := Acting{target: codeWoodDoor}
	breakOpen := Acting{target: codeGlassDoor}
	keyOpen := Acting{target: codeGoalDoor}

	actMap = map[Act]Acting{
		upAct:        up,
		downAct:      down,
		rightAct:     right,
		leftAct:      left,
		openAct:      open,
		breakOpenAct: breakOpen,
		keyOpenAct:   keyOpen,
		// getHammer :
	}
}

func initActCommandMap() {
	actCommandMap = map[ActCommand][]Act{
		"위": {upAct},
		"앞": {upAct},
		"상": {upAct},
		"북": {upAct},

		"아래": {downAct},
		"밑":  {downAct},
		"하":  {downAct},
		"남":  {downAct},

		"오른": {rightAct},
		"우":  {rightAct},
		"동":  {rightAct},

		"왼": {leftAct},
		"좌": {leftAct},
		"서": {leftAct},

		"연":  {openAct, breakOpenAct, keyOpenAct},
		"열":  {openAct, breakOpenAct, keyOpenAct},
		"사용": {openAct, breakOpenAct, keyOpenAct},
		"이용": {openAct, breakOpenAct, keyOpenAct},

		"부수": {breakOpenAct},
		"부순": {breakOpenAct},
		"깨":  {breakOpenAct},
		"깬":  {breakOpenAct},

		// "줍": {getHammer, getKey},
	}
}

func (act Act) getActing() Acting {
	return actMap[act]
}

func move(act []Act) {
	actName := act[0]
	move := actName.getActing()
	if move.target == codeFloor {
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

		moveScript.print(actName)

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
