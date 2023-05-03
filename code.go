package main

import "fmt"

type Code int

const door = 100
const item = 10

const (
	codePlayer Code = iota - 1
	codeFloor
	codeBlank
)

const (
	codeKey Code = (iota + 1) * item
	codeHammer
	codeHand
)

const (
	codeCloseDoor Code = (iota + 1) * door
	codeGlassDoor
	codeWoodDoor
)

func (code Code) getDoorNumber() int {
	return int(code / door)
}

func (code Code) getItemNumber() int {
	return int(code / item)
}

func (code Code) isDoor() bool {
	if code.getDoorNumber() > 0 {
		return true
	}

	return false
}

func (code Code) isItem() bool {
	if code.getItemNumber() > 0 {
		return true
	}

	return false
}

func (code Code) isOpenDoor() bool {
	if code.isDoor() && code.isItem() && code.getDoorNumber() == code.getItemNumber() {
		return true
	}
	return false
}

func getAroundDoorCoords( /* codes ...*Code */ ) *Code {
	aroundPlace := []*Code{playInfo.upPlace, playInfo.downPlace, playInfo.rightPlace, playInfo.leftPlace}
	for _, place := range aroundPlace {
		fmt.Println("test : ", place)
		fmt.Println("*test : ", *place)
		fmt.Println()
		if (*place).isDoor() {
			return place
		}
	}

	return nil
}