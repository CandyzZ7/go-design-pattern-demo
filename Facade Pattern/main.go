package main

import "fmt"

// Hero represents a game hero with operations
type Hero struct {
	Name string
}

func (h *Hero) Introduce() {
	fmt.Println("你好，我是", h.Name)
}
func (h *Hero) BuyEquipment() {
	fmt.Println("购买装备多兰盾")
}

func (h *Hero) UpgradeSkills() {
	fmt.Println("升级E技能")
}

// GameFacade provides a simple interface for controlling a game hero
type GameFacade struct {
	hero *Hero
}

func (gf *GameFacade) PrepareForBattle() {
	gf.hero.Introduce()
	gf.hero.BuyEquipment()
	gf.hero.UpgradeSkills()
}

func main() {
	hero := &Hero{
		Name: "盖伦",
	}
	facade := &GameFacade{hero}
	facade.PrepareForBattle()
}
