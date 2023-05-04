package main

func initGame() {
	initField()
	initAttributeMap()
	initActMap()
	initPlayInfo(Coords{y: 4, x: 1}, Coords{y: 0, x: 7})
}
