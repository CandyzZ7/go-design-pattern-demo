package main

import "fmt"

// Hero 英雄：游戏中的角色
type Hero struct {
	Name  string
	State State // 维护当前状态
}

// State 状态：定义状态接口
type State interface {
	Handle() // 处理状态
}

// NormalState 正常状态：实现状态接口
type NormalState struct{}

func (ns *NormalState) Handle() {
	fmt.Println("英雄状态：正常")
}

// InjuredState 受伤状态：实现状态接口
type InjuredState struct{}

func (is *InjuredState) Handle() {
	fmt.Println("英雄状态：受伤")
}

func main() {
	// 创建英雄
	hero := &Hero{Name: "Ashe"}

	// 设置英雄状态为正常状态
	hero.State = &NormalState{}
	hero.State.Handle()

	// 当英雄受伤时，改变状态为受伤状态
	hero.State = &InjuredState{}
	hero.State.Handle()
}
