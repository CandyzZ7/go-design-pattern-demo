package main

import "fmt"

type Hero interface {
	Attack() string
}

type Warrior struct{}

func (w *Warrior) Attack() string {
	return "Attack with sword!"
}

type Marksman struct{}

func (m *Marksman) Shoot() string {
	return "Shoot with bow!"
}

type MarksmanAdapter struct {
	Marksman Marksman
}

func (ma *MarksmanAdapter) Attack() string {
	return ma.Marksman.Shoot()
}

func main() {
	warrior := &Warrior{} // 创建一个战士英雄

	// 输出战士英雄的攻击方式
	fmt.Println(warrior.Attack())

	marksman := &Marksman{} // 创建一个射手英雄
	// 创建一个 MarksmanAdapter，并接受一个射手英雄
	marksmanAdapter := &MarksmanAdapter{
		Marksman: *marksman,
	}

	// 输出使用适配器的射手英雄的攻击方式
	fmt.Println(marksmanAdapter.Attack())
}
