package game

var equipmentMap map[Component]Stat

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

func (component Component) getEquipmentInfo() Stat {
	return equipmentMap[component]
}

func (component Component) getOffense() int {
	return equipmentMap[component].offence
}

func (component Component) getDefense() int {
	return equipmentMap[component].defense
}
