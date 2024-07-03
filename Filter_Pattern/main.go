package main

import (
	"fmt"
)

// Champion 定义英雄联盟角色
type Champion struct {
	Name   string
	Role   string
	Gender string
}

// FilterFunc 过滤函数类型
type FilterFunc func(champion Champion) bool

// RoleFilter 根据角色过滤
func RoleFilter(role string) FilterFunc {
	return func(champion Champion) bool {
		return champion.Role == role
	}
}

// GenderFilter 根据性别过滤
func GenderFilter(gender string) FilterFunc {
	return func(champion Champion) bool {
		return champion.Gender == gender
	}
}

// ApplyFilters 应用过滤器链
func ApplyFilters(filters []FilterFunc, champions []Champion) []Champion {
	var filteredChampions []Champion

	for _, champion := range champions {
		matched := true
		for _, filter := range filters {
			if !filter(champion) {
				matched = false
				break
			}
		}
		if matched {
			filteredChampions = append(filteredChampions, champion)
		}
	}

	return filteredChampions
}

// PrintChampions 打印角色列表
func PrintChampions(champions []Champion) {
	for _, champion := range champions {
		fmt.Println(" -", champion.Name)
	}
}

func main() {
	champions := []Champion{
		{Name: "Ahri", Role: "Mage", Gender: "Female"},
		{Name: "Garen", Role: "Fighter", Gender: "Male"},
		{Name: "Jinx", Role: "Marksman", Gender: "Female"},
		{Name: "Ashe", Role: "Marksman", Gender: "Female"},
		{Name: "Braum", Role: "Support", Gender: "Male"},
	}

	// 定义过滤器链
	filters := []FilterFunc{
		RoleFilter("Mage"),
		GenderFilter("Female"),
	}

	// 应用过滤器链并打印结果
	filteredChampions := ApplyFilters(filters, champions)
	fmt.Println("Mages or Females:")
	PrintChampions(filteredChampions)

	// 重置过滤器链
	filters = []FilterFunc{
		RoleFilter("Support"),
		GenderFilter("Male"),
	}

	// 应用过滤器链并打印结果
	filteredChampions = ApplyFilters(filters, champions)
	fmt.Println("\nMale Supports:")
	PrintChampions(filteredChampions)
}
