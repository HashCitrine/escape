package game

import "strings"

type Command string

type Acting struct {
	target Code
	coords Coords
}

var movementCommandMap map[Command]Movement
var interactionCommandMap map[Command][]Interaction

func initMovementCommandMap() {
	movementCommandMap = map[Command]Movement{
		"위": codeUp,
		"앞": codeUp,
		"상": codeUp,
		"북": codeUp,

		"아래": codeDown,
		"밑":  codeDown,
		"하":  codeDown,
		"남":  codeDown,

		"오른": codeRight,
		"우":  codeRight,
		"동":  codeRight,

		"왼": codeLeft,
		"좌": codeLeft,
		"서": codeLeft,

		/* "연":  {codeOpen, codeBreak, codeUnlock},
		"열":  {codeOpen, codeBreak, codeUnlock},
		"사용": {codeOpen, codeBreak, codeUnlock},
		"이용": {codeOpen, codeBreak, codeUnlock},

		"부수": {codeBreak},
		"부순": {codeBreak},
		"깨":  {codeBreak},
		"깬":  {codeBreak}, */

		// "줍": {getHammer, getKey},
	}
}

func initInteractionCommandMap() {
	interactionCommandMap = map[Command][]Interaction{
		"연":  {codeOpen, codeBreak, codeUnlock},
		"열":  {codeOpen, codeBreak, codeUnlock},
		"사용": {codeOpen, codeBreak, codeUnlock},
		"이용": {codeOpen, codeBreak, codeUnlock},

		"부수": {codeBreak},
		"부순": {codeBreak},
		"깨":  {codeBreak},
		"깬":  {codeBreak},

		"줍": {codeGet},

		"공격": {codeAttack},
		"도망": {codeRun},
	}
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