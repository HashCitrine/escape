package game

func (component Component) getEquipmentInfo() Stat {
	return equipmentMap[component]
}

func (component Component) getOffense() int {
	return equipmentMap[component].offence
}

func (component Component) getDefense() int {
	return equipmentMap[component].defense
}
