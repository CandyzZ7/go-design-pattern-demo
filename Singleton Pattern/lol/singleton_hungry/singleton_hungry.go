package main

import (
	"fmt"
)

type SummonersRift struct {
	Heroes []string // 存储召唤者峡谷中的所有英雄
}

var riftInstance *SummonersRift

func GetRiftInstance() *SummonersRift {
	return riftInstance
}

func init() {
	riftInstance = &SummonersRift{} // 饿汉式，直接初始化实例
}

func main() {
	// 召唤者峡谷1
	rift1 := GetRiftInstance()
	rift1.Heroes = append(rift1.Heroes, "盖伦") // 添加盖伦到英雄列表

	// 召唤者峡谷2
	rift2 := GetRiftInstance()
	rift2.Heroes = append(rift2.Heroes, "艾希") // 在召唤者峡谷2里添加艾希

	// 召唤者峡谷1
	fmt.Println(rift1.Heroes) // 输出: [盖伦 艾希]，因为rift1和rift2实际上是同一个实例
	// 召唤者峡谷2
	fmt.Println(rift2.Heroes) // 输出: [盖伦 艾希]，因为rift1和rift2实际上是同一个实例

}
