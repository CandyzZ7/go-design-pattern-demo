package main

import (
	"fmt"
)

// Hero 结构体，包含名字、角色、武器和特技
type Hero struct {
	Name        string
	Role        string
	Weapon      string
	SpecialMove string
}

// Builder 接口，包含了创建英雄所需的步骤
type Builder interface {
	SetName(string) Builder
	SetRole(string) Builder
	SetWeapon(string) Builder
	SetSpecialMove(string) Builder
	Build() *Hero
}

// HeroBuilder 通用的建造者，包含通用的构建步骤
type HeroBuilder struct {
	hero *Hero
}

func (hb *HeroBuilder) SetName(name string) Builder {
	hb.hero.Name = name
	return hb
}

func (hb *HeroBuilder) SetRole(role string) Builder {
	hb.hero.Role = role
	return hb
}

func (hb *HeroBuilder) SetWeapon(weapon string) Builder {
	hb.hero.Weapon = weapon
	return hb
}

func (hb *HeroBuilder) SetSpecialMove(specialMove string) Builder {
	hb.hero.SpecialMove = specialMove
	return hb
}

func (hb *HeroBuilder) Build() *Hero {
	return hb.hero
}

// GalenBuilder 盖伦的建造者，复用 HeroBuilder 的步骤
type GalenBuilder struct {
	HeroBuilder
}

func NewGalenBuilder() *GalenBuilder {
	return &GalenBuilder{HeroBuilder: HeroBuilder{hero: &Hero{}}}
}

// AsheBuilder 艾希的建造者，复用 HeroBuilder 的步骤
type AsheBuilder struct {
	HeroBuilder
}

func NewAsheBuilder() *AsheBuilder {
	return &AsheBuilder{HeroBuilder: HeroBuilder{hero: &Hero{}}}
}

func main() {
	// 创建盖伦
	gb := NewGalenBuilder()
	galen := gb.SetName("盖伦").SetRole("战士").SetWeapon("大剑").SetSpecialMove("大旋风").Build()
	fmt.Println(galen)

	// 创建艾希
	ab := NewAsheBuilder()
	ashe := ab.SetName("艾希").SetRole("射手").SetWeapon("冰弓").SetSpecialMove("魔极冰射").Build()
	fmt.Println(ashe)
}
