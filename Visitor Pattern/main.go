package main

import "fmt"

// Hero 结构定义了游戏中的一个单位，包括其名字
type Hero struct {
	Name string
}

// GetName 是一个给定英雄的方法，用于返回其名字
func (h Hero) GetName() string {
	// 返回英雄名
	return h.Name
}

// Visitor 结构定义了一个游戏访客，可以查看但不能更改游戏数据
type Visitor struct {
	Heroes []Hero
}

// VisitHeroes 会依次访问并打印出访客看到的所有英雄的名字
func (v Visitor) VisitHeroes() {
	for _, hero := range v.Heroes {
		fmt.Println(hero.GetName())
	}
}

func main() {
	// 创建一些示例英雄
	heroes := []Hero{
		{"Galen"},
		{"Teemo"},
		{"Ashe"},
	}

	// 创建一个访客，并让其访问所有英雄
	visitor := Visitor{Heroes: heroes}
	visitor.VisitHeroes()
}
