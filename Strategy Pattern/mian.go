package main

import "fmt"

type OutfitStrategy interface {
	Outfit()
}

type HappyOutfit struct{}

func (h *HappyOutfit) Outfit() {
	fmt.Println("S12亚索快乐流已配置----Wegame一键天赋已配置")
}

type CritOutfit struct{}

func (k *CritOutfit) Outfit() {
	fmt.Println("S12亚索暴击流已配置----Wegame一键天赋已配置")
}

type Hero struct {
	name     string
	strategy OutfitStrategy
}

func (h *Hero) SetStrategy(strategy OutfitStrategy) {
	h.strategy = strategy
}
func (h *Hero) UseOutfit() {
	h.strategy.Outfit()
}

func main() {
	hero := &Hero{
		name: "亚索",
	}
	hero.SetStrategy(&HappyOutfit{})
	hero.UseOutfit()
	hero.SetStrategy(&CritOutfit{})
	hero.UseOutfit()
}
