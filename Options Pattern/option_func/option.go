package main

import "fmt"

type Hero struct {
	Name     string
	Level    int
	Position string
}

type Option func(*Hero)

func WithName(name string) Option {
	return func(h *Hero) {
		h.Name = name
	}
}

func WithLevel(level int) Option {
	return func(h *Hero) {
		h.Level = level
	}
}

func WithPosition(position string) Option {
	return func(h *Hero) {
		h.Position = position
	}
}

// NewHero 生成一个新的英雄
func NewHero(opts ...Option) *Hero {
	h := &Hero{} // 创建一个新的英雄
	for _, opt := range opts {
		opt(h) // 应用所有的选项
	}
	return h
}

func main() {
	hero := NewHero(WithName("亚索"), WithLevel(15), WithPosition("中路"))
	fmt.Printf("Hero: %v\n", hero)
}
