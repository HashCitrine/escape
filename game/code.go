package game

type Code int
type Codetype string
type Movement Code
type Interaction Code

const (
	field       Codetype = "field"
	item        Codetype = "item"
	door        Codetype = "door"
	movement    Codetype = "movement"
	interaction Codetype = "interaction"
	enemy       Codetype = "enemy"
	box 		Codetype = "box"
)

const (
	codeFloor = iota + 1
	codeBlank
)

const (
	codeGoalDoor Code = iota + 1
	codeGlassDoor
	codeWoodDoor
)

const (
	codeKey    = codeGoalDoor
	codeHammer = codeGlassDoor
	codeHand   = codeWoodDoor
	codePortion = iota + 1 + codeHand
	codeWoodSword
	codeIronSword
	codeWoodShield
	codeLeatherRobe
	codeLeatherPants
	codeLeatherHat
)

const (
	codeUp Movement = iota + 1
	codeDown
	codeRight
	codeLeft
)

const (
	codeOpen Interaction = iota + 1
	codeBreak
	codeUnlock
	codeGet
	codeAttack
	codeRun
	// codeClose
	// codeLock
)

const (
	codeSquirrel = iota + 1
	codeRabbit
	codeDeer
)

const (
	codeOpenBox = iota + 1
	codeCloseBox
)