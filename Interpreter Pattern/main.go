package main

import (
	"fmt"
)

// RuneExpression 定义符文表达式接口
type RuneExpression interface {
	Interpret()
}

// AttackRuneExpression 实现攻击符文表达式
type AttackRuneExpression struct{}

func (re *AttackRuneExpression) Interpret() {
	fmt.Println("增加10%的攻击力")
}

// SpeedRuneExpression 实现移动速度符文表达式
type SpeedRuneExpression struct{}

func (re *SpeedRuneExpression) Interpret() {
	fmt.Println("增加15%的移动速度")
}

// PlayerRunePage 玩家选择的符文编译
type PlayerRunePage struct {
	playerRunes []RuneExpression
}

func (prp *PlayerRunePage) addRuneExpression(re RuneExpression) {
	prp.playerRunes = append(prp.playerRunes, re)
}

func (prp *PlayerRunePage) interpret() {
	for _, re := range prp.playerRunes {
		re.Interpret()
	}
}

// 使用这些表达式
func main() {
	playerRunePage := &PlayerRunePage{}
	playerRunePage.addRuneExpression(&AttackRuneExpression{})
	playerRunePage.addRuneExpression(&SpeedRuneExpression{})

	// 解释并执行
	playerRunePage.interpret()
}
