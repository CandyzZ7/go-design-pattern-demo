package main

import "fmt"

// Hero 是我们的英雄接口，它定义了一系列的游戏准备步骤
type Hero interface {
	Prepare()
	BuyItems()
	UpgradeSkills()
	ReadyToGame()
}

// LenientMode 是我们的模板方法，它是 Hero 游戏准备的模板
func LenientMode(h Hero) {
	h.Prepare()
	h.BuyItems()
	h.UpgradeSkills()
	h.ReadyToGame()
}

// Mage 是我们的英雄：法师
type Mage struct{}

func (m *Mage) Prepare() {
	fmt.Println("法师准备进入游戏...")
}

// BuyItems 法师的装备购买
func (m *Mage) BuyItems() {
	fmt.Println("法师购买法杖...")
}

func (m *Mage) UpgradeSkills() {
	fmt.Println("法师升级火球术技能...")
}

func (m *Mage) ReadyToGame() {
	fmt.Println("法师准备好进入游戏！")
}

// Marksman 是我们的英雄：射手
type Marksman struct{}

func (m *Marksman) Prepare() {
	fmt.Println("射手准备进入游戏...")
}

// 射手的装备购买
func (m *Marksman) BuyItems() {
	fmt.Println("射手购买短弓...")
}

func (m *Marksman) UpgradeSkills() {
	fmt.Println("射手升级多重箭技能...")
}

func (m *Marksman) ReadyToGame() {
	fmt.Println("射手准备好进入游戏！")
}

func main() {
	mage := &Mage{}         // 创建一个法师英雄
	marksman := &Marksman{} // 创建一个射手英雄

	fmt.Println("========== 法师的游戏准备 ==========")
	LenientMode(mage)

	fmt.Println("\n========== 射手的游戏准备 ==========")
	LenientMode(marksman)
}
