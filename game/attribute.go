package game

import "fmt"

type Attribute struct {
	commands []string
	mark     string
	place    []*Code
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
