package main

import (
	"fmt"
	"sync"
)

type SummonersRift struct {
	Heroes []string // 存储召唤者峡谷中的所有英雄
}

var (
	riftInstance *SummonersRift
	once         sync.Once
)

func GetRiftInstance() *SummonersRift {
	once.Do(func() {
		riftInstance = &SummonersRift{} // 延迟初始化，如果没有获取实例的需求，就不会创建实例
	})
	return riftInstance
}

func main() {
	// 召唤者峡谷1
	rift1 := GetRiftInstance()
	rift1.Heroes = append(rift1.Heroes, "盖伦") // 添加盖伦到英雄列表

	// 召唤者峡谷2
	rift2 := GetRiftInstance()
	rift2.Heroes = append(rift2.Heroes, "艾希") // 在召唤者峡谷2里添加Ashe

	// 召唤者峡谷1
	fmt.Println(rift1.Heroes) // 输出: [盖伦 艾希]，因为rift1和rift2实际上是同一个实例
	// 召唤者峡谷2
	fmt.Println(rift2.Heroes) // 输出: [盖伦]，因为rift1和rift2实际上是同一个实例

}
