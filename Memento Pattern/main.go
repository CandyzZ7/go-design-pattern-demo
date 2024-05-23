package main

import "fmt"

// Hero 英雄：游戏中的角色
type Hero struct {
	Name   string
	Health int
	Mana   int
}

// GameState 游戏状态：保存英雄的状态
type GameState struct {
	Health int
	Mana   int
}

// GameStateManager 游戏状态管理器：负责管理游戏状态
type GameStateManager struct {
	states []*GameState
}

// SaveState 保存状态：保存英雄的当前状态
func (g *GameStateManager) SaveState(hero *Hero) {
	g.states = append(g.states, &GameState{hero.Health, hero.Mana})
}

// RestoreState 恢复状态：恢复英雄到之前保存的状态
func (g *GameStateManager) RestoreState(hero *Hero, index int) {
	if index >= 0 && index < len(g.states) {
		state := g.states[index]
		hero.Health = state.Health
		hero.Mana = state.Mana
	}
}

func main() {
	// 创建英雄
	hero := &Hero{Name: "Ashe", Health: 100, Mana: 50}

	// 创建游戏状态管理器
	stateManager := &GameStateManager{}

	// 保存当前状态
	stateManager.SaveState(hero)

	// 模拟游戏过程：英雄受到伤害和恢复
	fmt.Println("初始状态:", hero)
	hero.Health -= 30
	hero.Mana -= 10
	fmt.Println("受到伤害后:", hero)

	// 再次保存当前状态
	stateManager.SaveState(hero)

	// 恢复到初始状态
	stateManager.RestoreState(hero, 0)
	fmt.Println("恢复后:", hero)
}
