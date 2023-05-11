package game

// type Act string
type ActCommand string

type Acting struct {
	target Code
	coords Coords
}

// var actMap map[Act]Acting
var movingCommandMap map[ActCommand]Moving
var actionCommandMap map[ActCommand][]Actioning

/* const (
	upAct        Act = "위"
	downAct      Act = "아래"
	rightAct     Act = "오른쪽"
	leftAct      Act = "왼쪽"
	openAct      Act = "열기"
	breakOpenAct Act = "부숴서열기"
	keyOpenAct   Act = "열쇠열기"
) */

/* func initActMap() {
	upCoords, downCoords, rightCoords, leftCoords := getAroundCoords(Coords{})
	up := Acting{target: codeFloor, coords: upCoords}
	down := Acting{target: codeFloor, coords: downCoords}
	right := Acting{target: codeFloor, coords: rightCoords}
	left := Acting{target: codeFloor, coords: leftCoords}

	open := Acting{target: codeWoodDoor}
	breakOpen := Acting{target: codeGlassDoor}
	keyOpen := Acting{target: codeGoalDoor}

	actMap = map[Act]Acting{
		codeUp:       up,
		codeDown:      down,
		rightAct:     right,
		leftAct:      left,
		openAct:      open,
		breakOpenAct: breakOpen,
		keyOpenAct:   keyOpen,
		// getHammer :
	}
} */

func initMovingCommandMap() {
	movingCommandMap = map[ActCommand]Moving{
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
	actionCommandMap = map[ActCommand][]Actioning{
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

/* func (act Act) getActing() Acting {
	return actMap[act]
} */

func (act Moving) getDirectionInfo() Coords {
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

func (act Moving) getDirectionName() string {
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
