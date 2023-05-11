package game

type Code int

type Act Code
type Moving Act
type Actioning Act

const item = 10
const door = 100
const act = 1000
const moving = act
const action = act

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

const (
	codeUp Moving = (iota + 1) * moving
	codeDown
	codeRight
	codeLeft
)

const (
	codeOpen Actioning = Actioning(codeLeft) + (iota + 1) * action
	codeBreak
	codeUnlock
	codeClose
	codeLock
	codeGet
)

func (code Code) getActNumber() int {
	return int(code) / act
}

func (code Code) getDoorNumber() int {
	return int(code) % act / door
}

func (code Code) getItemNumber() int {
	return int(code) % door / item
}

func (code Code) isActioning() bool {
	if code.getActNumber() > 0 {
		return true
	}

	return false
}

func (code Code) isDoor() bool {
	if code.getActNumber() == 0 && code.getDoorNumber() < 10 && code.getDoorNumber() > 0 {
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

/* func checkActToDoor(actArray []Act, doorCode Code, ifDoor *Code, ifDoorIsOpen bool) bool {
	if actArray != nil {
		for _, act := range actArray {
			if act.getActing().target == doorCode {
				return true
			}
		}
	}

	if ifDoorIsOpen {
		alreadyOpenDoorScript.print((*ifDoor).getName())
	}

	doNotActToDoorScript.print((*ifDoor).getName())

	return false
} */

func (code Code) getName() string {
	return attributeMap[code].getName()
}

func (actioning Code) isCanActioning(door Code, item Code) bool {
	if door.getDoorNumber() == item.getItemNumber() && item.getItemNumber() == actioning.getActNumber() - 4 {
		return true
	}

	return false
}