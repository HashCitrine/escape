package main

type Act string

type Acting struct {
	direction Code
	coords    Coords
}

var actMap map[Act][]Acting

func initActMap() {
	upCoords,downCoords,rightCoords,leftCoords := getAroundCoords(Coords{})
	up := Acting{direction: codeFloor, coords: upCoords}
	down := Acting{direction: codeFloor, coords: downCoords}
	right := Acting{direction: codeFloor, coords: rightCoords}
	left := Acting{direction: codeFloor, coords: leftCoords}

	open := Acting{direction: codeWoodDoor}
	breakOpen := Acting{direction: codeGlassDoor}
	keyOpen := Acting{direction: codeCloseDoor}

	actMap = map[Act][]Acting{
		"위": {up},
		"앞": {up},
		"상": {up},
		"북": {up},

		"아래": {down},
		"밑":  {down},
		"하":  {down},
		"남":  {down},

		"오른": {right},
		"우":  {right},
		"동":  {right},

		"왼": {left},
		"좌": {left},
		"서": {left},

		"연": {open, breakOpen, keyOpen},
		"열": {open, breakOpen, keyOpen},

		"부수" : {breakOpen},
		"부순" : {breakOpen},
		"깨" : {breakOpen},
		"깬" : {breakOpen},
	}
}