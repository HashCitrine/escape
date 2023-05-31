package game

type Block struct {
	parts    []Component
	passable bool
}

func (block Block) isDoor() bool {
	return len(block.parts) == 1 && block.parts[0].isDoor()
}

func (block *Block) isOpen() bool {
	return (*block).passable && (*block).isDoor()
}

func (block Block) getDoorName() string {
	if block.isDoor() {
		return block.parts[0].getName()
	}

	return ""
}

func (block Block) getDoor() Component {
	if block.isDoor() {
		return block.parts[0]
	}

	return Component{}
}

func (block *Block) findItem() []Component {
	var result []Component
	parts := (*block).parts

	if len(parts) > 0 {
		for _, component := range parts {
			if component.codetype == item {
				result = append(result, component)
			}
		}
	}

	return result
}

func (block *Block) printFloorItems() {
	parts := block.parts

	for _, part := range parts {
		findSomethingScript.print(part.getName())
	}
}
