package main

import (
	"fmt"
)

type Hero struct {
	name string
}

func (h *Hero) UseSkills() {
	fmt.Println(h.name, "使用技能")
}

func (h *Hero) Move() {
	fmt.Println(h.name, "移动")
}

type Command interface {
	CommandControl()
}

type Summoner struct {
	name string
	cmd  []Command
}

func (s *Summoner) AddCommand(c Command) {
	s.cmd = append(s.cmd, c)
}

func (s *Summoner) RemoveCommand(c Command) {
	for i, v := range s.cmd {
		if v == c {
			s.cmd = append(s.cmd[:i], s.cmd[i+1:]...)
		}
	}
}

func (s *Summoner) Notify() {
	for _, v := range s.cmd {
		fmt.Println(s.name, "发出命令")
		v.CommandControl()
	}
}

type CommandUseSkills struct {
	hero *Hero
}

func (c *CommandUseSkills) CommandControl() {
	c.hero.UseSkills()
}

type CommandMove struct {
	hero *Hero
}

func (c *CommandMove) CommandControl() {
	c.hero.Move()
}
func main() {
	s := &Summoner{
		name: "BLGYagao",
	}
	s.AddCommand(&CommandUseSkills{hero: &Hero{name: "亚索"}})
	s.AddCommand(&CommandMove{hero: &Hero{name: "亚索"}})
	s.Notify()
}
