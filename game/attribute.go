package game

import "fmt"

type Attribute struct {
	commands []string
	icon     string
	place    []*Component
}

var attributeMap map[Component]Attribute

const (
	playerIcon = "🐈"
	floorIcon  = "⬜"
	blankIcon  = "⬛"
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
	key := getAttribute(getStringArray("열쇠", "키"), "🗝️", getPlace(5, 5))
	hammer := getAttribute(getStringArray("망치", "해머", "함마"), "🔨", getPlace(2, 0))
	hand := getAttribute(getStringArray("손"), "", nil)

	openGoalDoor := getAttribute(goalDoor.commands, "%", nil)
	openGlassDoor := getAttribute(glassDoor.commands, "≠", nil)
	openWoodDoor := getAttribute(woodDoor.commands, "○", nil)

	attributeMap = map[Component]Attribute{
		door.getComponent(codeGoalDoor, false):  goalDoor,
		door.getComponent(codeGlassDoor, false): glassDoor,
		door.getComponent(codeWoodDoor, false):  woodDoor,

		item.getComponent(codeKey, true):    key,
		item.getComponent(codeHammer, true): hammer,
		item.getComponent(codeHand, true):   hand,

		door.getComponent(codeGoalDoor, true):  openGoalDoor,
		door.getComponent(codeGlassDoor, true): openGlassDoor,
		door.getComponent(codeWoodDoor, true):  openWoodDoor,
	}
}

func getIcon(icon interface{}) string {
	return fmt.Sprintf("[%v]", icon)
}

func getStringArray(stringArray ...string) []string {
	return stringArray
}

func getAttribute(commands []string, icon string, place ...*Component) Attribute {
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
				*place = component
			}
		}
	}
}

func (attribute Attribute) getPlace() []*Component {
	return attribute.place
}

func (attribute Attribute) getName() string {
	return attribute.commands[0]
}
