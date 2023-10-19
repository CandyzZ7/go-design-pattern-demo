package main

import "fmt"

// Party 是参与战斗的各方（英雄或小兵）的接口
type Party interface {
	ChangeGold(value int)
}

// Mediator 是中介者接口，用于协调英雄和小兵的交互
type Mediator interface {
	HeroBoughtEquipment(value int)
	MinionBeKilled(value int)
}

// Hero 的结构
type Hero struct {
	Mediator Mediator // 持有中介者的对象，用于后续的通信
	Gold     int      // 英雄的金币数量
}

// ChangeGold 英雄可以改变自己的金币数
func (h *Hero) ChangeGold(value int) {
	h.Gold += value
}

// BuyEquipment 英雄可以通过中介者购买装备
func (h *Hero) BuyEquipment(value int) {
	h.Mediator.HeroBoughtEquipment(value)
}

// Minion 的结构
type Minion struct {
	Mediator Mediator // 持有中介者的对象，用于后续的通信
	Gold     int      // 小兵携带的金币
}

// BeKilled 小兵可以通过中介者被杀
func (m *Minion) BeKilled() {
	m.Mediator.MinionBeKilled(m.Gold)
}

// Artifact 是实现了 Mediator 接口的中介者，是一个遗物
type Artifact struct {
	Hero   *Hero
	Minion *Minion
}

// HeroBoughtEquipment 英雄购买装备时调用
func (a *Artifact) HeroBoughtEquipment(value int) {
	a.Hero.ChangeGold(-value)
	fmt.Printf("Hero bought an equipment for %d gold.\n", value)
}

// MinionBeKilled 英雄杀死小兵时调用
func (a *Artifact) MinionBeKilled(value int) {
	a.Hero.ChangeGold(value)
	fmt.Printf("Hero got %d gold for killing a minion.\n", value)
}

// 主函数模拟游戏逻辑
func main() {
	hero := &Hero{Gold: 100}                          // 初始化一个英雄，初始金币为100
	minion := &Minion{Gold: 10}                       // 初始化一个小兵，携带金币为10
	artifact := &Artifact{Hero: hero, Minion: minion} // 初始化一个遗物
	hero.Mediator = artifact                          // 分别把遗物设为英雄和小兵的中介者
	minion.Mediator = artifact

	fmt.Printf("Hero initial gold: %d\n", hero.Gold)

	minion.BeKilled()     // 英雄击杀小兵，会得到小兵身上的全部金币
	hero.BuyEquipment(30) // 英雄花费30金币购买装备

	fmt.Printf("Hero final gold: %d\n", hero.Gold) // 最后英雄剩余的金币为：初始金币100 + 小兵身上的金币10 - 购买装备的金币30 = 80
}
