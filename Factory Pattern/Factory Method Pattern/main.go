package main

import "fmt"

// IHero 是所有英雄的通用接口
type IHero interface {
	Play()
}

// IHeroFactory 是所有英雄工厂的通用接口
type IHeroFactory interface {
	CreateHero() IHero
}

// Galen 表示一个具体的英雄
type Galen struct{}

// Play 方法实现了 Hero 接口要求的方法
func (g *Galen) Play() {
	fmt.Println("正在使用盖伦英雄")
}

// GalenFactory 是盖伦工厂
type GalenFactory struct{}

// CreateHero 方法创建一个盖伦英雄，实现了 HeroFactory 接口要求的方法
func (gf *GalenFactory) CreateHero() IHero {
	return &Galen{}
}

// Ashe 表示另一个具体的英雄
type Ashe struct{}

// Play 方法实现了 Hero 接口要求的方法
func (a *Ashe) Play() {
	fmt.Println("正在使用艾希英雄")
}

// AsheFactory 是艾希工厂
type AsheFactory struct{}

// CreateHero 方法创建一个艾希英雄，实现了 HeroFactory 接口要求的方法
func (af *AsheFactory) CreateHero() IHero {
	return &Ashe{}
}

func main() {
	var factory IHeroFactory

	// 玩家选择盖伦
	factory = &GalenFactory{}
	hero := factory.CreateHero()
	hero.Play()

	// 玩家选择艾希
	factory = &AsheFactory{}
	hero = factory.CreateHero()
	hero.Play()
}
