package main

import "fmt"

// ChampionFactory 英雄工厂
type ChampionFactory struct {
	champions map[string]Champion
}

// NewChampionFactory 创建英雄工厂
func NewChampionFactory() *ChampionFactory {
	return &ChampionFactory{
		champions: make(map[string]Champion),
	}
}

// GetChampion 获取英雄
func (f *ChampionFactory) GetChampion(name string) Champion {
	if champion, ok := f.champions[name]; ok {
		return champion
	}
	champion := NewConcreteChampion(name)
	f.champions[name] = champion
	return champion
}

// Champion 英雄接口
type Champion interface {
	Select() // 选择英雄
}

// ConcreteChampion 具体英雄
type ConcreteChampion struct {
	Name string
	// 在这里可以添加更多的属性
}

// NewConcreteChampion 创建具体英雄
func NewConcreteChampion(name string) Champion {
	return &ConcreteChampion{Name: name}
}

// Select 选择英雄
func (c *ConcreteChampion) Select() {
	fmt.Printf("选择了英雄：%s\n", c.Name)
}

func main() {
	factory := NewChampionFactory()

	// 选择英雄
	hero1 := factory.GetChampion("Ashe")
	hero1.Select()

	hero2 := factory.GetChampion("Teemo")
	hero2.Select()

	hero3 := factory.GetChampion("Ashe")
	hero3.Select()

	// 输出：
	// 选择了英雄：Ashe
	// 选择了英雄：Teemo
	// 选择了英雄：Ashe（来自享元模式，不需要重新创建）
}
