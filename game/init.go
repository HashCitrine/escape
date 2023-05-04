package game

import (
	. "escape/game/util"
)

var fieldArray [][]Code

var attributeMap map[Code]Attribute

var actMap map[Act]Acting
var actCommandMap map[ActCommand][]Act

var playInfo PlayInfo

func initGame() {
	initField()
	initAttributeMap()
	initActMap()
	initPlayInfo(Coords{Y: 4, X: 1}, Coords{Y: 0, X: 7})
}

func initField() {
	fieldArray = make([][]Code, 6)

	for i := range fieldArray {
		fieldArray[i] = make([]Code, 8)
		for j := range fieldArray[i] {
			fieldArray[i][j] = codeBlank
		}
	}
}

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
	glassDoor := getAttribute(getStringArray("유리문", "유리", "하늘"), "=", getPlace(2, 2))
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

func initActMap() {
	upCoords, downCoords, rightCoords, leftCoords := GetAroundCoords(Coords{})
	up := Acting{name: "위", direction: codeFloor, coords: upCoords}
	down := Acting{name: "아래", direction: codeFloor, coords: downCoords}
	right := Acting{name: "오른쪽", direction: codeFloor, coords: rightCoords}
	left := Acting{name: "왼쪽", direction: codeFloor, coords: leftCoords}

	open := Acting{direction: codeWoodDoor}
	breakOpen := Acting{direction: codeGlassDoor}
	keyOpen := Acting{direction: codeGoalDoor}

	actMap = map[Act]Acting{
		upAct:        up,
		downAct:      down,
		rightAct:     right,
		leftAct:      left,
		openAct:      open,
		breakOpenAct: breakOpen,
		keyOpenAct:   keyOpen,
	}

	actCommandMap = map[ActCommand][]Act{
		"위": {upAct},
		"앞": {upAct},
		"상": {upAct},
		"북": {upAct},

		"아래": {downAct},
		"밑":  {downAct},
		"하":  {downAct},
		"남":  {downAct},

		"오른": {rightAct},
		"우":  {rightAct},
		"동":  {rightAct},

		"왼": {leftAct},
		"좌": {leftAct},
		"서": {leftAct},

		"연":  {openAct, breakOpenAct, keyOpenAct},
		"열":  {openAct, breakOpenAct, keyOpenAct},
		"사용": {openAct, breakOpenAct, keyOpenAct},
		"이용": {openAct, breakOpenAct, keyOpenAct},

		"부수": {breakOpenAct},
		"부순": {breakOpenAct},
		"깨":  {breakOpenAct},
		"깬":  {breakOpenAct},
	}
}

func initPlayInfo(currentCoords Coords, goalCoords Coords) {
	playInfo = PlayInfo{
		goalCoords:    goalCoords,
		currentCoords: currentCoords,
		inventory:     []Code{codeHand},
	}

	setPlayInfo(currentCoords)
	return
}
