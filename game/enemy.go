package game

type Stat struct {
	offence int
	defense int
}

type Charactor struct {
	component Component
	maxHp     int
	hp        int
	common    Stat
}

type Enemy Charactor

