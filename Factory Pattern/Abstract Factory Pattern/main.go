package main

import "fmt"

// Hero 定义英雄和符文的接口
type Hero interface {
	Fight()
}
type Runes interface {
	Activate()
}

// Galen 创建盖伦英雄
type Galen struct{}

func (Galen) Fight() { fmt.Println("盖伦正在战斗") }

// AttackRunes 创建攻击符文
type AttackRunes struct{}

func (AttackRunes) Activate() { fmt.Println("攻击符文被激活了") }

// Ashe 创建艾希英雄
type Ashe struct{}

func (Ashe) Fight() { fmt.Println("艾希正在战斗") }

// SpeedRunes 创建速度符文
type SpeedRunes struct{}

func (SpeedRunes) Activate() { fmt.Println("速度符文被激活了") }

// Factory 抽象工厂接口
type Factory interface {
	GetHero() Hero
	GetRunes() Runes
}

// GalenFactory 盖伦工厂，创建盖伦和攻击符文
type GalenFactory struct{}

func (GalenFactory) GetHero() Hero   { return &Galen{} }
func (GalenFactory) GetRunes() Runes { return &AttackRunes{} }

// AsheFactory 艾希工厂，创建艾希和速度符文
type AsheFactory struct{}

func (AsheFactory) GetHero() Hero   { return &Ashe{} }
func (AsheFactory) GetRunes() Runes { return &SpeedRunes{} }

func main() {
	var factory Factory

	// 创建盖伦和攻击符文
	factory = GalenFactory{}
	hero := factory.GetHero()
	hero.Fight()
	runes := factory.GetRunes()
	runes.Activate()

	// 创建艾希和速度符文
	factory = AsheFactory{}
	hero = factory.GetHero()
	hero.Fight()
	runes = factory.GetRunes()
	runes.Activate()
}
