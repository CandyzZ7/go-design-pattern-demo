package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// GenerateChampions 生成大量的英雄数据
func GenerateChampions(numChampions int) []Champion {
	var champions []Champion
	for i := 0; i < numChampions; i++ {
		champion := Champion{
			Name:   fmt.Sprintf("Champion%d", i),
			Role:   randomRole(),
			Gender: randomGender(),
		}
		champions = append(champions, champion)
	}
	return champions
}

func randomRole() string {
	roles := []string{"Mage", "Fighter", "Marksman", "Support"}
	return roles[rand.Intn(len(roles))]
}

func randomGender() string {
	genders := []string{"Male", "Female"}
	return genders[rand.Intn(len(genders))]
}

// BenchmarkApplyFilters benchmarks ApplyFilters using serial execution
func BenchmarkApplyFilters(b *testing.B) {
	numChampions := 1000000
	champions := GenerateChampions(numChampions)

	filters := []FilterFunc{
		RoleFilter("Mage"),
		GenderFilter("Female"),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ApplyFilters(filters, champions)
	}
}

// BenchmarkApplyFiltersConcurrent benchmarks ApplyFiltersConcurrent using concurrent goroutines
func BenchmarkApplyFiltersConcurrent(b *testing.B) {
	numChampions := 1000000
	champions := GenerateChampions(numChampions)

	filters := []FilterFunc{
		RoleFilter("Mage"),
		GenderFilter("Female"),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ApplyFiltersConcurrent(filters, champions)
	}
}
