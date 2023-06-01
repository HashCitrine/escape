package game

func init() {
	initMovementCommandMap()
	initInteractionCommandMap()
	initEquipmentMap()
	initDropItemMap()
	initEnenyMap()
}

var movementCommandMap map[Command]Movement
var interactionCommandMap map[Command][]Interaction
var equipmentMap map[Component]Stat
var dropItemMap map[Component][]DropItem
var enemyMap map[*Block]Enemy

func initMovementCommandMap() {
	movementCommandMap = map[Command]Movement{
		"위": codeUp,
		"앞": codeUp,
		"상": codeUp,
		"북": codeUp,

		"아래": codeDown,
		"밑":  codeDown,
		"하":  codeDown,
		"남":  codeDown,

		"오른": codeRight,
		"우":  codeRight,
		"동":  codeRight,

		"왼": codeLeft,
		"좌": codeLeft,
		"서": codeLeft,

		/* "연":  {codeOpen, codeBreak, codeUnlock},
		"열":  {codeOpen, codeBreak, codeUnlock},
		"사용": {codeOpen, codeBreak, codeUnlock},
		"이용": {codeOpen, codeBreak, codeUnlock},

		"부수": {codeBreak},
		"부순": {codeBreak},
		"깨":  {codeBreak},
		"깬":  {codeBreak}, */

		// "줍": {getHammer, getKey},
	}
}

func initInteractionCommandMap() {
	interactionCommandMap = map[Command][]Interaction{
		"연":  {codeOpen, codeBreak, codeUnlock},
		"열":  {codeOpen, codeBreak, codeUnlock},
		"사용": {codeOpen, codeBreak, codeUnlock},
		"이용": {codeOpen, codeBreak, codeUnlock},

		"부수": {codeBreak},
		"부순": {codeBreak},
		"깨":  {codeBreak},
		"깬":  {codeBreak},

		"줍": {codeGet},

		"도망": {codeRun},
		"공격": {codeAttack},
		"방어": {codeShield},
		"회복": {codeRecovery},

		"착용": {codeWear},
		"장착": {codeWear},
		"장비": {codeWear},
		"입기": {codeWear},
	}
}

func initEquipmentMap() {
	woodSword := Stat{offence: 5, defense: 0}
	ironSword := Stat{offence: 10, defense: 0}
	woodShield := Stat{offence: 0, defense: 10}
	leatherRobe := Stat{offence: 0, defense: 6}
	leatherPants := Stat{offence: 0, defense: 4}
	leatherShoes := Stat{offence: 0, defense: 3}

	equipmentMap = map[Component]Stat{
		item.getComponent(codeWoodSword):    woodSword,
		item.getComponent(codeIronSword):    ironSword,
		item.getComponent(codeWoodShield):   woodShield,
		item.getComponent(codeLeatherRobe):  leatherRobe,
		item.getComponent(codeLeatherPants): leatherPants,
		item.getComponent(codeLeatherShoes): leatherShoes,
	}
}

func initDropItemMap() {
	portionItem := item.getComponent(codePortion)
	woodSwordItem := item.getComponent(codeWoodSword)
	ironSwordItem := item.getComponent(codeIronSword)
	// woodShield := item.getComponent(codeWoodShield)
	leatherRobeItem := item.getComponent(codeLeatherRobe)
	leatherPantsItem := item.getComponent(codeLeatherPants)
	leatherHatItem := item.getComponent(codeLeatherShoes)

	portion1 := DropItem{item: portionItem, amount: 1, probability: 15}
	portion2 := DropItem{item: portionItem, amount: 2, probability: 10}
	portion3 := DropItem{item: portionItem, amount: 3, probability: 5}
	woodSword := DropItem{item: woodSwordItem, amount: 1, probability: 20}
	ironSword := DropItem{item: ironSwordItem, amount: 1, probability: 15}
	leatherRobe := DropItem{item: leatherRobeItem, amount: 1, probability: 7}
	leatherPants := DropItem{item: leatherPantsItem, amount: 1, probability: 8}
	leatherHat := DropItem{item: leatherHatItem, amount: 1, probability: 10}
	notDrop := DropItem{probability: 10}
	boxDrop := []DropItem{
		notDrop, woodSword, ironSword,
		leatherRobe, leatherPants, leatherHat,
		portion1, portion2, portion3,
	}

	portion1 = DropItem{item: portionItem, amount: 1, probability: 70}
	notDrop = DropItem{probability: 30}
	squirrelDrop := []DropItem{
		portion1, notDrop,
	}

	portion1 = DropItem{item: portionItem, amount: 1, probability: 50}
	portion2 = DropItem{item: portionItem, amount: 2, probability: 30}
	notDrop = DropItem{probability: 20}
	rabbitDrop := []DropItem{
		portion1, portion2, notDrop,
	}

	key := DropItem{item: item.getComponent(codeKey), amount: 1, probability: 100}
	deerDrop := []DropItem{
		key,
	}

	dropItemMap = map[Component][]DropItem{
		box.getComponent(codeCloseBox):         boxDrop,
		enemy.getComponent(Code(codeSquirrel)): squirrelDrop,
		enemy.getComponent(Code(codeRabbit)):   rabbitDrop,
		enemy.getComponent(Code(codeDeer)):     deerDrop,
	}
}

func initEnenyMap() {
	squirrel := Enemy{
		component: Component{code: Code(codeSquirrel), codetype: enemy},
		hp:        50,
		common: Stat{
			offence: 5,
			defense: 0,
		},
	}

	rabbit := Enemy{
		component: Component{code: Code(codeRabbit), codetype: enemy},
		hp:        70,
		common: Stat{
			offence: 7,
			defense: 3,
		},
	}

	deer := Enemy{
		component: Component{code: Code(codeDeer), codetype: enemy},
		hp:        100,
		common: Stat{
			offence: 10,
			defense: 5,
		},
	}

	enemyMap = map[*Block]Enemy{
		getPlace(4, 8): squirrel,
		getPlace(3, 6): squirrel,

		getPlace(0, 0): rabbit,
		getPlace(7, 6): rabbit,

		getPlace(0, 0): deer,
		getPlace(7, 2): deer,
	}
}