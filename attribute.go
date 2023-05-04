package main

import "fmt"

type Attribute struct {
	commands []string
	mark     string
	place    []*Code
}

var attributeMap map[Code]Attribute

func initAttributeMap() {
	// Structor
	start := getAttribute(getStringArray("시작 지점"), "1", getPlace(4, 1))
	floor := getAttribute(getStringArray("복도"), " ",
		getPlace(0, 4), getPlace(0, 5), getPlace(0, 6),
		getPlace(2, 1), getPlace(2, 3), getPlace(2, 4),
		getPlace(3, 1), getPlace(3, 4),
		getPlace(4, 4), getPlace(4, 5))
	blank := getAttribute(getStringArray("공백"), "", nil)
	/* blank := GetCommon(GetStringArray("공백"), "",
	GetPlace(0, 0), GetPlace(0, 1), GetPlace(0, 2), GetPlace(0, 3),
	GetPlace(1, 0), GetPlace(1, 1), GetPlace(1, 2), GetPlace(1, 3), GetPlace(1, 5), GetPlace(1, 6), GetPlace(1, 7),
	GetPlace(2, 5), GetPlace(2, 6), GetPlace(2, 7),
	GetPlace(3, 0), GetPlace(3, 2), GetPlace(3, 3), GetPlace(3, 5), GetPlace(3, 6), GetPlace(3, 7),
	GetPlace(4, 0), GetPlace(4, 2), GetPlace(4, 3), GetPlace(4, 6), GetPlace(4, 7),
	GetPlace(5, 0), GetPlace(5, 1), GetPlace(5, 2), GetPlace(5, 3), GetPlace(5, 4), GetPlace(5, 6), GetPlace(5, 7)) */

	// Door
	goalDoor := getAttribute(getStringArray("회색문", "회색"), "&", getPlace(0, 7))
	glassDoor := getAttribute(getStringArray("유리문", "유리", "하늘",), "=", getPlace(2, 2))
	woodDoor := getAttribute(getStringArray("나무문", "나무", "갈색"), "◐", getPlace(1, 4))

	// Item
	key := getAttribute(getStringArray("열쇠", "키"), "K", getPlace(5, 5))
	hammer := getAttribute(getStringArray("망치", "해머", "오함마"), "H", getPlace(2, 0))
	hand := getAttribute(getStringArray("손"), "", nil)

	openGoalDoor := getAttribute(goalDoor.commands, "%", nil)
	openGlassDoor := getAttribute(glassDoor.commands, "≠", nil)
	openWoodDoor := getAttribute(woodDoor.commands, "○", nil)

	attributeMap = map[Code]Attribute{
		codePlayer: start,
		codeFloor:  floor,
		codeBlank:  blank,

		codeGoalDoor:  goalDoor,
		codeGlassDoor: glassDoor,
		codeWoodDoor:  woodDoor,

		codeKey:    key,
		codeHammer: hammer,
		codeHand:   hand,

		codeGoalDoor + codeKey:     openGoalDoor,
		codeGlassDoor + codeHammer: openGlassDoor,
		codeWoodDoor + codeHand:    openWoodDoor,
	}

	setAttributeToField()
}

func getMark(mark interface{}) string {
	return fmt.Sprintf("[%v]", mark)
}

func getStringArray(stringArray ...string) []string {
	return stringArray
}


func getAttribute(commands []string, mark string, place ...*Code) Attribute {
	if mark != "" {
		mark = getMark(mark)
	}
	return Attribute{
		commands,
		mark,
		place,
	}
}

func setAttributeToField() {
	for attributeCode, attribute := range attributeMap {
		placeArr := attribute.place
		for _, place := range placeArr {
			if place != nil {
				*place = attributeCode
			}
		}
	}
}

func (attribute Attribute) getName() string {
	return attribute.commands[0]
}