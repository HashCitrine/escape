package game

import "fmt"

type Attribute struct {
	commands []string
	icon     string
	place    []*Code
}

func getIcon(icon interface{}) string {
	return fmt.Sprintf("[%v]", icon)
}

func getStringArray(stringArray ...string) []string {
	return stringArray
}

func getAttribute(commands []string, icon string, place ...*Code) Attribute {
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

