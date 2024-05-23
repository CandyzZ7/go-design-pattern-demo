package main

import "fmt"

// Champion 英雄：游戏中的角色
type Champion interface {
	Attack() // 攻击
}

// Equipment 装备：游戏中的道具
type Equipment interface {
	Use() // 使用
}

// ConcreteChampion 具体英雄
type ConcreteChampion struct {
	Name      string
	Equipment Equipment // 英雄装备
}

// Attack 攻击
func (c *ConcreteChampion) Attack() {
	fmt.Printf("%s 发动了攻击\n", c.Name)
	// 英雄使用装备
	c.Equipment.Use()
}

// Sword 剑：一种装备
type Sword struct{}

// Use 使用剑
func (s *Sword) Use() {
	fmt.Println("使用了剑进行攻击")
}

// Bow 弓箭：一种装备
type Bow struct{}

// Use 使用弓箭
func (b *Bow) Use() {
	fmt.Println("使用了弓箭进行攻击")
}

func main() {
	// 创建具体英雄
	hero1 := &ConcreteChampion{Name: "Ashe", Equipment: &Sword{}}
	hero2 := &ConcreteChampion{Name: "Teemo", Equipment: &Bow{}}

	// 英雄攻击
	hero1.Attack()
	hero2.Attack()
}
