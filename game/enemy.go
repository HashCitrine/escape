package game

type Stat struct {
	offence int
	defense int
}

type Charactor struct {
	component Component
	maxHp     int
	hp        int
	common    Stat
}

type Enemy Charactor

var squirrel = Enemy{
	component: Component{code: Code(codeSquirrel), codetype: enemy},
	hp:        50,
	common: Stat{
		offence: 5,
		defense: 0,
	},
}

var rabbit = Enemy{
	component: Component{code: Code(codeRabbit), codetype: enemy},
	hp:        70,
	common: Stat{
		offence: 7,
		defense: 3,
	},
}

var deer = Enemy{
	component: Component{code: Code(codeDeer), codetype: enemy},
	hp:        100,
	common: Stat{
		offence: 10,
		defense: 5,
	},
}

var enemyMap map[*Block]Enemy

func initEnenyMap() {
	enemyMap = map[*Block]Enemy{
		getPlace(4, 8): squirrel,
		getPlace(3, 6): squirrel,

		getPlace(0, 0): rabbit,
		getPlace(7, 6): rabbit,

		getPlace(0, 0): deer,
		getPlace(7, 2): deer,
	}
}

/* const (
	commandRun      Command = "도망"
	commandAttack   Command = "공격"
	commandShield   Command = "방어"
	commandRecovery Command = "회복"
) */

func (player Player) getDefense() int {
	return player.common.defense + player.head.getDefense() + player.top.getDefense() + player.pants.getDefense()
}
