package game

import "fmt"

type Script string
type Narration Script
type Fnarration Narration

const (
	clearConsole Script = "\033[H\033[2J"
	doNotTooManyActScript Script = "욕심 부리지 말자. 차근차근 하나씩 행동해야 한다."
	blankScript Script = "막힌 길이다. 다시 생각해보자."
)

const (
	startScript Narration = "당신은 미로의 함정에 빠졌습니다. 이곳을 빠져나가야 합니다."
	endScript   Narration = "밝은 빛이 보입니다. 당신은 탈출에 성공했습니다."
	questionScript    Narration = "어떤 행동을 하시겠습니까?"
)

const (
	lookAtTheDoorscript Fnarration = "%s에 %s이 있습니다."
	passDoorScript Fnarration = "%s을 지나왔다.\n"
	closeDoorScript Fnarration = "%s이 닫혀 있다. 이대로는 나아갈 수 없다.\n"
	moveScript Fnarration = "%s로 이동했다.\n"
	findItmeScript Fnarration = "%s가 떨어져 있다. 어딘가에 사용할 수 있을 것 같다. 챙겨놓도록 하자.\n"
	notHaveItemScript Fnarration = "%s를 가지고 있지 않다. 다른 방법을 찾아보자.\n"
	alreadyOpenDoor Fnarration = "%s은 이미 열려있다. 지나갈 수 있을 것 같다.\n"
	useItemToDoor Fnarration = "%s(으)로 %s을 열었다. 이제 지나갈 수 있다.\n"
	doNotUseItemToDoor Fnarration = "%s(으)로는 %s을 열 수 없다. 다른 방법을 찾아보자.\n"
	doNotActToDoor Fnarration = "%s은 굳게 닫혀있다.\n"
)

func (script Script) print() {
	fmt.Println(script)
	fmt.Println()
}

func (script Narration) print() {
	fmt.Printf("(%s)", script)
	fmt.Println()
}

func (script Fnarration) print(s ...any) {
	fmt.Printf(string(script), s...)
	fmt.Println()
}