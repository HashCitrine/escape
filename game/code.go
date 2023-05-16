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
	// codeClose
	// codeLock
	// codeGet
)