package main

import "fmt"

// Hero 定义英雄接口，每个英雄都应具有战斗的能力
type Hero interface {
	Fight()
}

// Galen 创建实际英雄类：盖伦
type Galen struct{}

func (Galen) Fight() {
	fmt.Println("盖伦正在展开德玛西亚正义")
}

// Summoner 创建英雄的代理类：召唤师
type Summoner struct {
	hero Hero // 持有实际的英雄对象
}

// Fight summoner的Fight行为实现定义
func (s *Summoner) Fight() {
	fmt.Println("召唤师准备召唤英雄...")
	s.hero.Fight() // 在代理的实现中调用真实英雄的方法
	fmt.Println("英雄战斗结束，召唤师结束控制")
}

func main() {
	hero := &Galen{}                  // 创建实际英雄：盖伦
	summoner := &Summoner{hero: hero} // 创建代理：召唤师， 委托给盖伦英雄
	summoner.Fight()                  // 由召唤师代理进行战斗操作
}
