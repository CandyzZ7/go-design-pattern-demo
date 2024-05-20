package main

import "fmt"

type Hero struct {
	Name     string
	Level    int
	Position string
}

type Option interface {
	Apply(*Hero)
}

func WithName(name string) Option {
	return withName{name}
}

type withName struct{ name string }

func (w withName) Apply(h *Hero) {
	h.Name = w.name
}

func WithLevel(level int) Option {
	return withLevel{level}
}

type withLevel struct{ level int }

func (w withLevel) Apply(h *Hero) {
	h.Level = w.level
}

func WithPosition(position string) Option {
	return withPosition{position}
}

type withPosition struct{ position string }

func (w withPosition) Apply(h *Hero) {
	h.Position = w.position
}

// NewHero 生成一个新的英雄
func NewHero(opts ...Option) *Hero {
	h := &Hero{}
	for _, opt := range opts {
		opt.Apply(h)
	}
	return h
}

func main() {
	hero := NewHero(WithName("亚索"), WithLevel(15), WithPosition("中路"))
	fmt.Printf("Hero: %v\n", hero)
}
