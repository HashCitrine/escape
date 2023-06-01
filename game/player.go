package game

import "fmt"

type Player struct {
	goalCoords    Coords
	name          string
	currentCoords Coords
	inventory     []Component
	head      Component
	top       Component
	pants     Component
	shoes     Component
	rightHand Component
	leftHand  Component
	fight	  bool
	Charactor
}

var player Player

func init() {
	initPlayer(Coords{y: 11, x: 12}, Coords{y: 2, x: 0})
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

func updatePlayerPlace(directionCoords Coords /* , tempPlace *Code */) {
	player.currentCoords = directionCoords
	getPlaceByCoords(directionCoords).printFloorItems()
}

func (playInfo Player) getAroundDoorCoords() (*Block, string) {
	coordsArray := getAroundCoords(playInfo.currentCoords)

	for i, coords := range coordsArray {
		place := getPlaceByCoords(coords)
		if place != nil && place.isDoor() {
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

func printEquipment() {
	head := player.head
	top := player.top
	pants := player.pants
	shoes := player.shoes
	rightHand := player.rightHand
	leftHand := player.leftHand

	fmt.Println("[장비]")
	fmt.Println("머리 : ", head.getName())
	fmt.Println("상의 : ", top.getName())
	fmt.Println("하의 : ", pants.getName())
	fmt.Println("신발 : ", shoes.getName())
	fmt.Println("오른손 : ", rightHand.getName())
	fmt.Println("왼손 : ", leftHand.getName())
	fmt.Println()
}

func (player Player) getDefense() int {
	return player.common.defense + player.head.getDefense() + player.top.getDefense() + player.pants.getDefense()
}
