package game

var equipmentMap map[Component]TacticalCommon

func initEquipmentMap() {
	woodSword := TacticalCommon {offence: 5, defense: 0}
	ironSword := TacticalCommon {offence: 10, defense: 0}
	woodShield := TacticalCommon {offence: 0, defense: 10}
	leatherRobe := TacticalCommon {offence: 0, defense: 6}
	leatherPants := TacticalCommon {offence: 0, defense: 4}
	leatherHat := TacticalCommon {offence: 0, defense: 3}

	equipmentMap = map[Component]TacticalCommon {
		item.getComponent(codeWoodSword) : woodSword,
		item.getComponent(codeIronSword) : ironSword,
		item.getComponent(codeWoodShield) : woodShield,
		item.getComponent(codeLeatherRobe) : leatherRobe,
		item.getComponent(codeLeatherPants) : leatherPants,
		item.getComponent(codeLeatherHat) : leatherHat,
	}
}

func (component Component) getEquipmentInfo() TacticalCommon {
	return equipmentMap[component]
}

func (component Component) getOffense() int {
	return equipmentMap[component].offence
}

func (component Component) getDefense() int {
	return equipmentMap[component].defense
}