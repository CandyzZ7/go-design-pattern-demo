package main

import "fmt"

// IHero 定义英雄接口
type IHero interface {
	GetName() string // 获得英雄名字的方法
}

// Galen 定义盖伦（Galen）类型
type Galen struct {
}

// GetName 盖伦类型实现 Hero 接口的 GetName 方法
func (g Galen) GetName() string {
	return "盖伦"
}

// Yasuo 定义亚索（Yasuo）类型
type Yasuo struct {
}

// GetName 亚索类型实现 Hero 接口的 GetName 方法
func (y Yasuo) GetName() string {
	return "亚索"
}

// HeroFactory 定义英雄工厂函数，根据给定的英雄名，生成对应的英雄实例
func HeroFactory(heroName string) IHero {
	switch heroName {
	case "盖伦":
		return Galen{} // 如果英雄名是"盖伦"，创建并返回 Galen类型的实例
	case "亚索":
		return Yasuo{} // 如果英雄名是"亚索"，创建并返回 Yasuo类型的实例
	default:
		return nil // 如果英雄名不符合条件，返回nil
	}
}

func main() {
	galen := HeroFactory("盖伦")
	fmt.Println(galen.GetName()) // 输出: 盖伦

	yasuo := HeroFactory("亚索")
	fmt.Println(yasuo.GetName()) // 输出: 亚索
}
