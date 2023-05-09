package game

var fieldArray [][]Code

var attributeMap map[Code]Attribute



var playInfo PlayInfo

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
	start := getAttribute(getStringArray("시작 지점"), "🐈", getPlace(4, 1))
	floor := getAttribute(getStringArray("복도"), "⬜",
		getPlace(0, 4), getPlace(0, 5), getPlace(0, 6),
		getPlace(2, 1), getPlace(2, 3), getPlace(2, 4),
		getPlace(3, 1), getPlace(3, 4),
		getPlace(4, 4), getPlace(4, 5))
	blank := getAttribute(getStringArray("공백"), "⬛", nil)

	// Door
	goalDoor := getAttribute(getStringArray("회색문", "회색"), "🟪", getPlace(0, 7))
	glassDoor := getAttribute(getStringArray("유리문", "유리", "하늘"), "🟦", getPlace(2, 2))
	woodDoor := getAttribute(getStringArray("나무문", "나무", "갈색"), "🟫", getPlace(1, 4))

	// Item
	key := getAttribute(getStringArray("열쇠", "키"), "🗝️", getPlace(5, 5))
	hammer := getAttribute(getStringArray("망치", "해머", "함마"), "🔨", getPlace(2, 0))
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



func initPlayInfo(currentCoords Coords, goalCoords Coords) {
	playInfo = PlayInfo{
		goalCoords:    goalCoords,
		currentCoords: currentCoords,
		inventory:     []Code{codeHand},
	}

	setPlayInfo(currentCoords)
	return
}
