package main

import "fmt"

// HeroList 定义创建迭代器的接口
type HeroList interface {
	Iterator() HeroIterator
}

// HeroIterator 定义迭代器的接口
type HeroIterator interface {
	IsDone() bool
	Next() *Hero
}

// Hero 游戏中的角色
type Hero struct {
	Name   string
	Health int
	Mana   int
}

// HeroCollection 实现英雄列表的聚合器
type HeroCollection struct {
	heroes []*Hero
}

// NewHeroCollection 构造函数
func NewHeroCollection() *HeroCollection {
	return &HeroCollection{
		heroes: make([]*Hero, 0),
	}
}

// AddHero 向英雄集合中添加英雄
func (hc *HeroCollection) AddHero(hero *Hero) {
	hc.heroes = append(hc.heroes, hero)
}

// Iterator 实现聚合器接口，返回迭代器实例
func (hc *HeroCollection) Iterator() HeroIterator {
	return &HeroListIterator{
		heroes: hc.heroes,
		index:  0,
	}
}

// HeroListIterator 实现迭代器接口
type HeroListIterator struct {
	heroes []*Hero
	index  int
}

// IsDone 判断是否遍历完成
func (li *HeroListIterator) IsDone() bool {
	return li.index >= len(li.heroes)
}

// Next 获取下一个英雄
func (li *HeroListIterator) Next() *Hero {
	if !li.IsDone() {
		hero := li.heroes[li.index]
		li.index++
		return hero
	}
	return nil
}

func main() {
	// 创建英雄列表
	heroCollection := NewHeroCollection()
	heroCollection.AddHero(&Hero{Name: "艾希", Health: 100, Mana: 50})
	heroCollection.AddHero(&Hero{Name: "盖伦", Health: 80, Mana: 100})
	heroCollection.AddHero(&Hero{Name: "亚索", Health: 120, Mana: 0})

	// 创建英雄迭代器并遍历
	iterator := heroCollection.Iterator()
	for !iterator.IsDone() {
		hero := iterator.Next()
		fmt.Printf("英雄：%s，生命值：%d，法力值：%d\n", hero.Name, hero.Health, hero.Mana)
	}
}
