package game

type Code int

const door = 100
const item = 10

const (
	codePlayer Code = iota - 1
	codeFloor
	codeBlank
)

const (
	codeKey Code = (iota + 1) * item
	codeHammer
	codeHand
)

const (
	codeGoalDoor Code = (iota + 1) * door
	codeGlassDoor
	codeWoodDoor
)

func (code Code) getDoorNumber() int {
	return int(code) / door
}

func (code Code) getItemNumber() int {
	return int(code) % door / item
}

func (code Code) isDoor() bool {
	if code.getDoorNumber() > 0 {
		return true
	}

	return false
}

func (code Code) isItem() bool {
	if code.getDoorNumber() == 0 && code.getItemNumber() < 10 && code.getItemNumber() > 0 {
		return true
	}

	return false
}

func (code Code) isOpenDoor() bool {
	if code.getDoorNumber() == code.getItemNumber() && code.isDoor() {
		return true
	}
	return false
}

func checkInventory(item Code) bool {
	for _, haveItem := range playInfo.inventory {
		if item == haveItem {
			return true
		}
	}
	return false
}

func checkActToDoor(actArray []ActName, doorCode Code, ifDoor *Code, ifDoorIsOpen bool) bool {
	if actArray != nil {
		for _, act := range actArray {
			if act.getActing().targetCode == doorCode {
				return true
			}
		}
	}

	if ifDoorIsOpen {
		alreadyOpenDoor.print((*ifDoor).getName())
	}

	doNotActToDoor.print((*ifDoor).getName())

	return false
}

func (code Code) getName() string {
	return attributeMap[code].getName()
}
