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
	{0, 5},
	{1, 5},
	{2, 0}, /* {2, 1}, */ {2, 2}, {2, 5},
	{3, 2}, {3, 3}, /* {3, 4}, */ {3, 5}, {3, 6}, {3, 7}, {3, 8}, {3, 9}, /* {3, 10}, */ {3, 11},
	{4, 2}, {4, 5}, {4, 7},
	{5, 2}, {5, 5}, {5, 7},
	/* {6, 2}, */ {6, 5}, {6, 7},
	{7, 3}, {7, 3}, {7, 4}, {7, 5}, {7, 7},
	{8, 7}, {8, 8}, {8, 9}, {8, 10}, {8, 11},
	{9, 7}, {9, 11},
	/* {10, 7}, */ {10, 11},
	{11, 7}, {11, 11},
}

func initAttributeMap() {
	// Door
	goalDoor := getAttribute(getStringArray("회색문", "회색"), "&", getPlace(2, 1))
	glassDoor := getAttribute(getStringArray("유리문", "유리", "하늘"), "=", getPlace(6, 2), getPlace(3, 10))
	woodDoor := getAttribute(getStringArray("나무문", "나무", "갈색"), "◐", getPlace(3, 4), getPlace(10, 7))

	// Item
	key := getAttribute(getStringArray("열쇠", "키"), keyIcon, nil)
	hammer := getAttribute(getStringArray("망치", "해머", "함마"), hammerIcon, getPlace(0, 5), getPlace(8, 7))
	hand := getAttribute(getStringArray("손"), "", nil)
	woodSword := getAttribute(getStringArray("목검"), "", getPlace(8, 11))
	portion := getAttribute(getStringArray("회복약", "포션"), "", getPlace(3, 7))

	/* openGoalDoor := getAttribute(goalDoor.commands, "%", nil)
	openGlassDoor := getAttribute(glassDoor.commands, "≠", nil)
	openWoodDoor := getAttribute(woodDoor.commands, "○", nil) */

	squirrel := getAttribute(getStringArray("다람쥐", "람쥐", "쥐"), "", getPlace(4, 7), getPlace(3, 5))
	rabbit := getAttribute(getStringArray("토끼"), "", getPlace(7, 5))
	deer := getAttribute(getStringArray("사슴", "시슴"), "", getPlace(7, 2))

	boxFloor := getAttribute(getStringArray("상자", "박스"), "", getPlace(0, 5), getPlace(3, 3), getPlace(3, 11), getPlace(11, 7))

	attributeMap = map[Component]Attribute{
		door.getComponent(codeGoalDoor):  goalDoor,
		door.getComponent(codeGlassDoor): glassDoor,
		door.getComponent(codeWoodDoor):  woodDoor,

		item.getComponent(codeKey):    key,
		item.getComponent(codeHammer): hammer,
		item.getComponent(codeHand):   hand,
		item.getComponent(codeWoodSword): woodSword,
		item.getComponent(codePortion) : portion,

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
