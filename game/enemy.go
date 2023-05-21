package game

import (
	"math/rand"
	"time"
)

type TacticalCommon struct {
	offence int
	defense int
}

type Charactor struct {
	component Component
	hp        int
	common    TacticalCommon
}

type Player struct {
	charactor Charactor
	head      Component
	top       Component
	pants     Component
	rightHand Component
	leftHand  Component
}

type Enemy Charactor

var squirrel = Enemy{
	component: Component{code: Code(codeSquirrel), codetype: enemy},
	hp:        50,
	common: TacticalCommon{
		offence: 5,
		defense: 0,
	},
}

var rabbit = Enemy{
	component: Component{code: Code(codeRabbit), codetype: enemy},
	hp:        70,
	common: TacticalCommon{
		offence: 7,
		defense: 3,
	},
}

var deer = Enemy{
	component: Component{code: Code(codeDeer), codetype: enemy},
	hp:        100,
	common: TacticalCommon{
		offence: 10,
		defense: 5,
	},
}

var enemyMap map[*Block]Enemy

func initEnenyMap() {
	enemyMap = map[*Block]Enemy{
		getPlace(0, 0): squirrel,

		getPlace(0, 0): rabbit,

		getPlace(0, 0): deer,
	}
}

var player = Player{
	charactor: Charactor{
		component: Component{},
		hp:        50,
		common: TacticalCommon{
			offence: 3,
			defense: 0,
		},
	},
}

const (
	commandRun      Command = "도망"
	commandAttack   Command = "공격"
	commandShield   Command = "방어"
	commandRecovery Command = "회복"
)

func (block *Block) combat(scan string) bool {
	enemy := enemyMap[block]

	if enemy.component.isEmpty()  {
		// todo : 적이 없다. - script
		return false
	}

	enemyOffence := enemy.common.offence
	defense := player.getDefense()

	switch Command(scan) {
	case commandRun:
		rand.Seed(time.Now().UnixNano())
		result := rand.Float64() * 100

		if result > 0.5 {
			return false
		}
	case commandAttack:
		playerOffence := player.charactor.common.offence
		rightAttack := gameInfo.player.rightHand.getOffense()
		leftAttack := gameInfo.player.leftHand.getOffense()

		playerOffence += playerOffence + rightAttack + leftAttack
		enemy.hp -= playerOffence

		// todo : playerOffence 만큼 달았다. - script

		if enemy.hp <= 0 {
			// todo : enenmy가 죽었다. - script
			enemy.component.Drop()
			return true
		}
	case commandShield:
		defense += player.rightHand.getDefense() + player.leftHand.getDefense()
	case commandRecovery:
		if useItem(item.getComponent(codePortion)) {
			enemy.hp += 30
		}
	}

	enemyOffence -= defense

	if enemyOffence < 0 {
		enemyOffence = 0
	}

	player.charactor.hp -= enemyOffence
	// todo : enemyOffence 만큼 체력이 달았다. - script

	return true
}

func (player Player) getDefense() int {
	return player.charactor.common.defense + player.head.getDefense() + player.top.getDefense() + player.pants.getDefense()
}
