package main

import (
	"fmt"
	"time"
)

// Hero 定义一个英雄（Hero）和他的力量（Power）
type Hero struct {
	Name string
}

type Power struct {
	HeroName string
	Value    int
}

// calculatePower 函数计算一个英雄的力量
func calculatePower(hero Hero, c chan<- Power) {
	// 假设计算力量是一个耗时操作，这里我们用 Sleep 来模拟。
	time.Sleep(time.Second * 1)

	// 假设力量就是英雄名字的字符数量。
	power := len(hero.Name)

	// 将计算得到的结果发送到 channel 中。
	c <- Power{HeroName: hero.Name, Value: power}
}

func main() {
	heroes := []Hero{{Name: "盖伦"}, {Name: "艾希"}, {Name: "易大师"}}

	powerChan := make(chan Power, len(heroes))

	// 扇出：将每个英雄的力量计算任务分派给不同的 goroutine。
	// 这样，这些任务就可以并行执行，以提高执行效率。
	for _, hero := range heroes {
		go calculatePower(hero, powerChan)
	}

	// 扇入：收集所有英雄的力量，
	// 这样所有的结果就被收集到了一个地方，方便后续处理。
	powers := make(map[string]int)
	for i := 0; i < len(heroes); i++ {
		power := <-powerChan
		powers[power.HeroName] = power.Value
	}

	// 打印所有英雄及其对应的力量
	for hero, power := range powers {
		fmt.Printf("%s 的力量是 %d\n", hero, power)
	}
}
