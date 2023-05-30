package game

import (
	"fmt"
	"strings"
)

type Script string
type Narration Script
type Fnarration Script

const name = "{name}"

const (
	clearConsole              Script = "\033[H\033[2J"
	doNotTooManyCommandScript Script = "한 번에 하나씩 행동할 수 있습니다."
	needActAnythingScript     Script = "무엇이든 행동해야 합니다."
	openBoxScript             Script = "상자가 이미 열려있습니다."
	whatsYourNameScript		  Script = "당신의 고양이의 이름은 무엇인가요?"
)

const (
	startScript    Narration = "당신의 고양이 "+ name + "(이)가 미로의 함정에 빠졌습니다. 빠져나갈 수 있도록 지시를 내려야 합니다."
	endScript      Narration = "밝은 빛이 보입니다. 당신의 고양이는 탈출에 성공했습니다."
	deadScript Narration = "당신의 고양이 "+ name + "(이)가 더 이상 움직이지 않습니다. 당신은 동물학대범입니다."
	questionScript Narration = "어떤 행동을 하시겠습니까?"
	noDropScript Narration = "아이템이 떨어지지 않았습니다."
)

const (
	blankScript               Fnarration = "%s은 막힌 길입니다."
	lookAtTheDoorscript       Fnarration = "%s에 %s이 있습니다."
	passDoorScript            Fnarration = name + "(이)가 %s을 지나왔습니다.\n"
	closeDoorScript           Fnarration = "%s이 닫혀 있다. 이대로는 나아갈 수 없습니다.\n"
	moveScript                Fnarration = name + "(이)가 %s로 이동했습니다.\n"
	// findItmeScript            Fnarration = "%s(이)가 떨어져 있습니다.\n"
	getItemScript             Fnarration = name + "(이)가 %s을 챙겼습니다.\n"
	notHaveItemScript         Fnarration = "%s를 가지고 있지 않습니다.\n"
	alreadyOpenDoorScript     Fnarration = "%s은 이미 열려있습니다. 지나갈 수 있습니다.\n"
	useItemToDoorScript       Fnarration = "%s(으)로 %s을 열었습니다. 이제 지나갈 수 있습니다.\n"
	canNotUseItemToDoorScript Fnarration = "%s(으)로는 %s을 열 수 없습니다.\n"
	doNotActToDoorScript      Fnarration = "%s은 굳게 닫혀있습니다.\n"
	doNotActByItemScript      Fnarration = "%s를 가지고 있습니다."
	// doNotActByDoorScript      Fnarration = "%s을 열 방법을 찾아야 한다."
	canNotFindAroundDoor Fnarration = "%s에 %s를 사용할 수 있을까? 확인해보자."
	dropItemScript       Fnarration = "%s를 떨어뜨렸다."
	findSomethingScript  Fnarration = "%s(이)가 있다."
	attackedScript Fnarration = "%s가 %d만큼 데미지를 입었다."
	enemyInfoScript Fnarration = "%s의 HP : %d"
	playerInfoScript Fnarration = "HP : %d"
	wearScript Fnarration = "%s를 착용했다."
)

func print(script string) {
	script = replaceName(script)
	fmt.Println(script)
	fmt.Println()
}

func (script Script) print() {
	tempScript := replaceName(string(script))
	print(tempScript)
}

func (script Narration) print() {
	tempScript := replaceName(string(script))
	print(fmt.Sprintf("(%s)", tempScript))
}

func (script Fnarration) print(s ...any) {
	tempScript := replaceName(string(script))
	print(fmt.Sprintf(string(tempScript), s...))
}

func replaceName(script string) string {
	return strings.Replace(script, name, player.name, 1)
}