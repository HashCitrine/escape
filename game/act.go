package game

type ActCommand string

type Acting struct {
	target Code
	coords Coords
}

var movingCommandMap map[ActCommand]Movement
var actionCommandMap map[ActCommand][]Interaction

func initMovingCommandMap() {
	movingCommandMap = map[ActCommand]Movement{
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

func initActionCommandMap() {
	actionCommandMap = map[ActCommand][]Interaction{
		"연":  {codeOpen, codeBreak, codeUnlock},
		"열":  {codeOpen, codeBreak, codeUnlock},
		"사용": {codeOpen, codeBreak, codeUnlock},
		"이용": {codeOpen, codeBreak, codeUnlock},

		"부수": {codeBreak},
		"부순": {codeBreak},
		"깨":  {codeBreak},
		"깬":  {codeBreak},
	}
}

func (act Movement) getDirectionInfo() Coords {
	coords := Coords{}
	coordsArray := getAroundCoords(coords)

	switch act {
	case codeUp:
		coords = coordsArray[0]
	case codeDown:
		coords = coordsArray[1]
	case codeRight:
		coords = coordsArray[2]
	case codeLeft:
		coords = coordsArray[3]
	}
	return coords
}

func (act Movement) getDirectionName() string {
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
}
