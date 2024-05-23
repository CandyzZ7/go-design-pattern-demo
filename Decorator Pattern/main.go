package main

import "fmt"

// Champion 英雄：游戏中的角色
type Champion interface {
	UseSkill() // 使用技能
}

// ConcreteChampion 具体英雄
type ConcreteChampion struct {
	Name string
}

// UseSkill 使用技能
func (c *ConcreteChampion) UseSkill() {
	fmt.Printf("%s 使用了普通技能\n", c.Name)
}

// SkillDecorator 技能装饰器接口
type SkillDecorator interface {
	UseSkill() // 使用技能
}

// Skill1Decorator 技能1装饰器
type Skill1Decorator struct {
	Champion Champion
}

// UseSkill 使用技能
func (d *Skill1Decorator) UseSkill() {
	d.Champion.UseSkill()
	fmt.Printf("%s 使用了技能1\n", d.Champion.(*ConcreteChampion).Name)
}

// Skill2Decorator 技能2装饰器
type Skill2Decorator struct {
	Champion Champion
}

// UseSkill 使用技能
func (d *Skill2Decorator) UseSkill() {
	d.Champion.UseSkill()
	fmt.Printf("%s 使用了技能2\n", d.Champion.(*ConcreteChampion).Name)
}

func main() {
	// 创建具体英雄
	hero := &ConcreteChampion{Name: "Ashe"}

	// 使用普通技能
	hero.UseSkill()

	// 添加技能1装饰器
	decorator1 := &Skill1Decorator{Champion: hero}
	decorator1.UseSkill()

	// 添加技能2装饰器
	decorator2 := &Skill2Decorator{Champion: hero}
	decorator2.UseSkill()
}
