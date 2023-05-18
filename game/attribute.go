package game

import (
	"fmt"
	"strings"
)

type Character struct {
	hp        int
	attack    int
	defense   int
	equipment []Code
}

type Attribute struct {
	commands []string
	icon     string
	place    []*Block
}

var attributeMap map[Component]Attribute

const (
	playerIcon    = "🐈"
	floorIcon     = "⬜"
	blankIcon     = "⬛"
	doorIcon      = "🔷"
	openDoorIcon  = "🔵"
	somethingIcon = "🔶"
	keyIcon       = "🗝️"
	hammerIcon    = "🔨"
)

var floorPlace = []Coords{
	{0, 4}, {0, 5}, {0, 6},
	{2, 1}, {2, 3}, {2, 4},
	{3, 1}, {3, 4},
	{4, 1}, {4, 4}, {4, 5}}

func initAttributeMap() {
	// Door
	goalDoor := getAttribute(getStringArray("회색문", "회색"), "&", getPlace(0, 7))
	glassDoor := getAttribute(getStringArray("유리문", "유리", "하늘"), "=", getPlace(2, 2))
	woodDoor := getAttribute(getStringArray("나무문", "나무", "갈색"), "◐", getPlace(1, 4))

	// Item
	key := getAttribute(getStringArray("열쇠", "키"), keyIcon, getPlace(5, 5))
	hammer := getAttribute(getStringArray("망치", "해머", "함마"), hammerIcon, getPlace(2, 0))
	hand := getAttribute(getStringArray("손"), "", nil)

	/* openGoalDoor := getAttribute(goalDoor.commands, "%", nil)
	openGlassDoor := getAttribute(glassDoor.commands, "≠", nil)
	openWoodDoor := getAttribute(woodDoor.commands, "○", nil) */

	squirrel := getAttribute(getStringArray("다람쥐", "람쥐", "쥐"), "", getPlace(0, 0))
	rabbit := getAttribute(getStringArray("토끼"), "", getPlace(0, 1))
	deer := getAttribute(getStringArray("사슴", "시슴"), "", getPlace(0, 2))

	boxFloor := getAttribute(getStringArray("상자", "박스"), "", getPlace(0, 2))

	attributeMap = map[Component]Attribute{
		door.getComponent(codeGoalDoor):  goalDoor,
		door.getComponent(codeGlassDoor): glassDoor,
		door.getComponent(codeWoodDoor):  woodDoor,

		item.getComponent(codeKey):    key,
		item.getComponent(codeHammer): hammer,
		item.getComponent(codeHand):   hand,

		enemy.getComponent(codeSquirrel): squirrel,
		enemy.getComponent(codeRabbit):   rabbit,
		enemy.getComponent(codeDeer):     deer,

		/* door.getComponent(codeGoalDoor):  openGoalDoor,
		door.getComponent(codeGlassDoor): openGlassDoor,
		door.getComponent(codeWoodDoor):  openWoodDoor, */

		box.getComponent(codeCloseBox): boxFloor,
	}
}

func getIcon(icon interface{}) string {
	return fmt.Sprintf("[%v]", icon)
}

func getStringArray(stringArray ...string) []string {
	return stringArray
}

func getAttribute(commands []string, icon string, place ...*Block) Attribute {
	if icon != "" {
		// icon = getIcon(icon)
	}
	return Attribute{
		commands,
		icon,
		place,
	}
}

func setAttributeToField() {
	for component, attribute := range attributeMap {
		placeArr := attribute.getPlace()
		for _, place := range placeArr {
			if place != nil {
				componentArray := (*place).parts
				(*place).parts = append(componentArray, component)

				if !(*place).isDoor() && !(*place).passable {
					(*place).passable = true
				}
			}
		}
	}
}

func (attribute Attribute) getPlace() []*Block {
	return attribute.place
}

func (attribute Attribute) getName() string {
	return attribute.commands[0]
}

func GetDoorAndItem(scan string) (commandDoor Component, commandItem Component) {
	// var commandDoor Component
	// var commandItem Component
	// var commandEnemy Component

	for component, attribute := range attributeMap {
		/* if component.isOpen() {
			continue
		} */

		var tooMany bool
		for _, command := range attribute.commands {
			if strings.Contains(scan, command) {
				switch component.codetype {
				case door:
					tooMany, commandDoor = commandDoor.ifNotEmpty(component)
				case item:
					tooMany, commandItem = commandItem.ifNotEmpty(component)
					// case enemy:
					// 	tooMany, commandEnemy = commandEnemy.ifNotEmpty(component)
				}
			}

			if tooMany {
				doNotTooManyCommandScript.print()
				return
			}
		}
	}
	return
}
