package game

import "strings"

type Command string

type Acting struct {
	target Code
	coords Coords
}

func (act Movement) getDirectionInfo() (coords Coords, directionName string) {
	// coords := Coords{}
	coordsArray := getAroundCoords(coords)

	switch act {
	case codeUp:
		coords = coordsArray[0]
		directionName = "위쪽"
	case codeDown:
		coords = coordsArray[1]
		directionName = "아래쪽"
	case codeRight:
		coords = coordsArray[2]
		directionName = "오른쪽"
	case codeLeft:
		coords = coordsArray[3]
		directionName = "왼쪽"
	}
	// return coords

	return
}

/* func (act Movement) getDirectionName() string {
	var direction string
	switch act {
	case codeUp:
		direction = "위"
	case codeDown:
		direction = "아래"
	case codeRight:
		direction = "오른쪽"
	case codeLeft:
		direction = "왼쪽"
	}

	return direction
} */

func GetInteraction(scan string) (interactionArray []Interaction) {
	// var interaction Interaction
	for command, interactions := range interactionCommandMap {
		if strings.Contains(scan, string(command)) {
			interactionArray = interactions
		}
	}

	return
}