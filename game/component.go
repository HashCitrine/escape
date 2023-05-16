package game

type Component struct {
	code     Code
	codetype Codetype
	passable bool
}

func (codetype Codetype) getComponent(code Code, passable bool) Component {
	return Component{
		code:     code,
		codetype: codetype,
		passable: passable,
	}
}

/* func (component Component) getCode() Code {
	return component.code
} */

func (component Component) getName() string {
	return attributeMap[component].getName()
}

func (component Component) isOpen() bool {
	return component.codetype == door && component.passable == true
}

func (directionCoords Coords) isPassable() bool {
	if checkOutFieldByCoords(directionCoords) {
		return false
	}

	return getPlaceByCoords(directionCoords).passable
}

func (component Component) isItem() bool {
	return component.codetype == item
}

func (component Component) isDoor() bool {
	return component.codetype == door
}
func (component Component) isEmpty() bool {
	return component == Component{}
}

func (component Component) equals(tempComponent Component) bool {
	return component == tempComponent
}

func (tempDoor Component) tryOpenDoor(tempItem Component) bool {
	return tempDoor.codetype == door && tempItem.codetype == item && tempDoor.code == tempItem.code
}

func (interaction Interaction) isCanDo(door Component, item Component) bool {
	return Code(interaction) == door.code && door.code == item.code
}
