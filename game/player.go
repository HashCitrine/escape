package game

import "fmt"

type Player struct {
	goalCoords    Coords
	name          string
	currentCoords Coords
	inventory     []Component
	// player       Player
	head      Component
	top       Component
	pants     Component
	shoes     Component
	rightHand Component
	leftHand  Component
	Charactor
}

func initPlayer(currentCoords Coords, goalCoords Coords) {

	whatsYourNameScript.print()
	scan := Scan()

	player = Player{
		goalCoords:    goalCoords,
		currentCoords: currentCoords,
		inventory:     []Component{item.getComponent(codeHand)},
		Charactor: Charactor{
			component: Component{},
			hp:        50,
			maxHp:     50,
			common: Stat{
				offence: 3,
				defense: 0,
			},
		},
		name: scan,
	}
}

func updatePlayerPlace(tempPlaceCoords Coords /* , tempPlace *Code */) {
	player.currentCoords = tempPlaceCoords
}

func (playInfo Player) getAroundDoorCoords() (*Block, string) {
	coordsArray := getAroundCoords(playInfo.currentCoords)

	for i, coords := range coordsArray {
		place := getPlaceByCoords(coords)
		if place != nil && (*place).isDoor() {
			_, directionName := Movement(i + 1).getDirectionInfo()
			return place, directionName
		}
	}

	return nil, ""
}

func hasItem(item Component) bool {
	for _, hasItem := range player.inventory {
		if hasItem.equals(item) {
			return true
		}
	}
	return false
}

func printInventory() {
	inventorys := player.inventory

	if len(inventorys) < 2 {
		return
	}

	fmt.Print("소지품 : ")

	hand := item.getComponent(codeHand)
	for i, inventory := range inventorys {
		if i > 1 {
			fmt.Print(", ")
		}

		if inventory == hand {
			continue
		}

		fmt.Print(attributeMap[inventory].getName())
	}

	fmt.Println()
	fmt.Println()
}
