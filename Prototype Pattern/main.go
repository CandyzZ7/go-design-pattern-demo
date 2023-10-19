package main

import (
	"fmt"
)

// HeroPrototype 英雄原型接口定义复制自身的方法
type HeroPrototype interface {
	Clone() HeroPrototype
}

// LeeSin 实现一个具体英雄源：盲僧
type LeeSin struct {
	name   string
	weapon string
}

// Clone 盲僧的克隆方法
func (l *LeeSin) Clone() HeroPrototype {
	return &LeeSin{name: l.name, weapon: l.weapon}
}

// Riven 实现一个具体英雄源：瑞文
type Riven struct {
	name   string
	weapon string
}

// Clone 瑞文的克隆方法
func (r *Riven) Clone() HeroPrototype {
	return &Riven{name: r.name, weapon: r.weapon}
}

func main() {
	// 初始化一个盲僧原型
	originalLeeSin := &LeeSin{name: "盲僧", weapon: "拳套"}
	// 克隆一个盲僧
	clonedLeeSin := originalLeeSin.Clone()

	fmt.Printf("Original LeeSin: %v\n", originalLeeSin)
	fmt.Printf("Cloned LeeSin: %v\n", clonedLeeSin)

	// 初始化一个瑞文原型
	originalRiven := &Riven{name: "瑞文", weapon: "大剑"}
	// 克隆一个瑞文
	clonedRiven := originalRiven.Clone()

	fmt.Printf("Original Riven: %v\n", originalRiven)
	fmt.Printf("Cloned Riven: %v\n", clonedRiven)
}
