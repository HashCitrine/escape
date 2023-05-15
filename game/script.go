package game

import "fmt"

type Script string
type Narration Script
type Fnarration Script

const (
	clearConsole          Script = "\033[H\033[2J"
	doNotTooManyActScript Script = "한 번에 하나씩 행동할 수 있습니다."
	needActAnythingScript Script = "무엇이든 행동해야 합니다."
)

const (
	startScript    Narration = "당신의 고양이가 미로의 함정에 빠졌습니다. 빠져나갈 수 있도록 지시를 내려야 합니다."
	endScript      Narration = "밝은 빛이 보입니다. 당신의 고양이는 탈출에 성공했습니다."
	questionScript Narration = "어떤 행동을 하시겠습니까?"
)

const (
	blankScript           	  Fnarration = "%s은 막힌 길입니다."
	lookAtTheDoorscript       Fnarration = "%s에 %s이 있습니다."
	passDoorScript            Fnarration = "고양이가 %s을 지나왔습니다.\n"
	closeDoorScript           Fnarration = "%s이 닫혀 있다. 이대로는 나아갈 수 없습니다.\n"
	moveScript                Fnarration = "고양이가 %s로 이동했습니다.\n"
	findItmeScript            Fnarration = "%s가 떨어져 있습니다.\n"
	getItemScript             Fnarration = "고양이가 %s을 챙겼습니다.\n"
	notHaveItemScript         Fnarration = "%s를 가지고 있지 않습니다.\n"
	alreadyOpenDoorScript     Fnarration = "%s은 이미 열려있습니다. 지나갈 수 있습니다.\n"
	useItemToDoorScript       Fnarration = "%s(으)로 %s을 열었습니다. 이제 지나갈 수 있습니다.\n"
	canNotUseItemToDoorScript Fnarration = "%s(으)로는 %s을 열 수 없습니다.\n"
	doNotActToDoorScript      Fnarration = "%s은 굳게 닫혀있습니다.\n"
	doNotActByItemScript      Fnarration = "%s를 가지고 있습니다."
	// doNotActByDoorScript      Fnarration = "%s을 열 방법을 찾아야 한다."
	canNotFindAroundDoor 	  Fnarration = "%s에 %s를 사용할 수 있을까? 확인해보자."
)

func print(script string) {
	fmt.Println(script)
	fmt.Println()
}

func (script Script) print() {
	print(string(script))
}

func (script Narration) print() {
	print(fmt.Sprintf("(%s)", script))
}

func (script Fnarration) print(s ...any) {
	print(fmt.Sprintf(string(script), s...))
}
