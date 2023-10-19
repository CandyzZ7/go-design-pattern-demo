package main

import "fmt"

// Handler 处理器接口
type Handler interface {
	HandleRequest(request int)
	SetNext(handler Handler) Handler
}

// AbstractHandler 提供下一个操作链接的抽象处理器
type AbstractHandler struct {
	next Handler
}

// SetNext 设置下一个处理器并返回该处理器
func (h *AbstractHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

// ArmorHandler 护甲减伤处理器
type ArmorHandler struct {
	AbstractHandler
}

// HandleRequest 处理伤害请求，如果伤害超过20，则减少20的伤害，并传给下一个处理器
func (h *ArmorHandler) HandleRequest(request int) {
	if request > 20 {
		request -= 20
		fmt.Println("护甲减少了20点伤害，剩余伤害: ", request)
	} else {
		fmt.Println("护甲完全吸收了伤害，剩余伤害: 0")
		return
	}

	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// MagicResistanceHandler 魔法抗性处理器
type MagicResistanceHandler struct {
	AbstractHandler
}

// HandleRequest 处理伤害请求，如果伤害超过10，则减少10的伤害，并传给下一个处理器
func (h *MagicResistanceHandler) HandleRequest(request int) {
	if request > 10 {
		request -= 10
		fmt.Println("魔法抗性减少了10点伤害，剩余伤害: ", request)
	} else {
		fmt.Println("魔法抗性完全吸收了伤害，剩余伤害: 0")
		return
	}

	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// TrueDamageHandler 真实伤害处理器
type TrueDamageHandler struct {
	AbstractHandler
}

// HandleRequest 处理伤害请求，直接打出伤害
func (h *TrueDamageHandler) HandleRequest(request int) {
	fmt.Println("造成了真实伤害: ", request)
}

func main() {
	// 初始化处理器
	armorHandler := &ArmorHandler{}
	// 魔抗处理器
	magicResistanceHandler := &MagicResistanceHandler{}
	// 真伤处理器
	trueDamageHandler := &TrueDamageHandler{}

	// 设置处理器链
	armorHandler.SetNext(magicResistanceHandler).
		SetNext(trueDamageHandler)

	// 造成伤害
	dmg := 100
	fmt.Println("造成的伤害: ", dmg)
	armorHandler.HandleRequest(dmg)
}
