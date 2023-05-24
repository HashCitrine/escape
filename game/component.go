package game

type Component struct {
	code     Code
	codetype Codetype
}

func (codetype Codetype) getComponent(code Code) Component {
	return Component{
		code:     code,
		codetype: codetype,
	}
}

/* func (component Component) getCode() Code {
	return component.code
} */

func (component Component) getName() string {
	return attributeMap[component].getName()
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

func (component Component) isEnemy() bool {
	return component.codetype == enemy
}

func (component Component) isEmpty() bool {
	return component == Component{}
}

func (prevComponent Component) ifNotEmpty(postComponent Component) (tooMany bool, component Component) {
	empty := prevComponent.isEmpty()
	tooMany = !empty

	switch empty {
	case true:
		component = postComponent
	case false:
		component = prevComponent
	}

	return
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

func (component Component) wear() {
	inventory := gameInfo.inventory

	for _, item := range inventory {
		if item == component {
			noWear := Component{}
			useItem(component)
			wearScript.print(component.getName());

			switch component.code {
			case codeWoodSword, codeIronSword, codeWoodShield:
				if player.rightHand == noWear {
					player.rightHand = component
					return
				}

				if player.leftHand != noWear {
					gameInfo.inventory = append(inventory, player.leftHand)
				}

				player.leftHand = component
			case codeLeatherRobe:
				if player.top != noWear {
					gameInfo.inventory = append(inventory, player.top)
				}

				player.top = component
			case codeLeatherPants:
				if player.pants != noWear {
					gameInfo.inventory = append(inventory, player.pants)
				}
				player.pants = component
			case codeLeatherShoes:
				if player.shoes != noWear {
					gameInfo.inventory = append(inventory, player.shoes)
				}
				player.shoes = component
			}
		}
	}
}