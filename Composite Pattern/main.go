package main

import "fmt"

// Component 表示组合中的对象
type Component interface {
	GetName() string
	GetHealth() int
	GetMana() int
}

// Leaf 表示树中的叶子节点对象，叶子节点没有子节点
type Leaf struct {
	Name   string
	Health int
	Mana   int
}

// GetName 获取英雄名称
func (l *Leaf) GetName() string {
	return l.Name
}

// GetHealth 获取英雄生命值
func (l *Leaf) GetHealth() int {
	return l.Health
}

// GetMana 获取英雄法力值
func (l *Leaf) GetMana() int {
	return l.Mana
}

// Composite 表示组合中的容器对象，可以包含其他组件
type Composite struct {
	Name       string
	Health     int
	Mana       int
	components []Component
}

// GetName 获取英雄名称
func (c *Composite) GetName() string {
	return c.Name
}

// GetHealth 获取英雄生命值，组合对象的生命值为其所有子对象的生命值之和
func (c *Composite) GetHealth() int {
	totalHealth := c.Health
	for _, component := range c.components {
		totalHealth += component.GetHealth()
	}
	return totalHealth
}

// GetMana 获取英雄法力值，组合对象的法力值为其所有子对象的法力值之和
func (c *Composite) GetMana() int {
	totalMana := c.Mana
	for _, component := range c.components {
		totalMana += component.GetMana()
	}
	return totalMana
}

// Add 添加组件到容器中
func (c *Composite) Add(component Component) {
	c.components = append(c.components, component)
}

func main() {
	// 创建叶子节点（英雄）
	leaf1 := &Leaf{Name: "Ashe", Health: 100, Mana: 50}
	leaf2 := &Leaf{Name: "Teemo", Health: 80, Mana: 100}

	// 创建容器节点（队伍）
	composite := &Composite{Name: "Blue Team", Health: 0, Mana: 0}
	composite.Add(leaf1)
	composite.Add(leaf2)

	// 输出队伍的生命值和法力值
	fmt.Printf("队伍：%s，生命值：%d，法力值：%d\n", composite.GetName(), composite.GetHealth(), composite.GetMana())
}
