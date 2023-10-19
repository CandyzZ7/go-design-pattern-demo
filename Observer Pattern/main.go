package main

import "fmt"

type Event struct {
	Not     Notifier
	One     Listener
	Another Listener
	Msg     string
}

type Notifier interface {
	AddListener(l Listener)
	RemoveListener(l Listener)
	Notify(e *Event)
}

type Listener interface {
	OnFriendBeFight(e *Event)
}

type Hero struct {
	Name  string
	Party string
}

func (h *Hero) GetHero() string {
	return fmt.Sprintf("姓名：%s，阵营：%s", h.Name, h.Party)
}
func (h *Hero) OnFriendBeFight(e *Event) {
	if h.Name == e.One.(*Hero).Name {
		return
	}
	if h.Name == e.Another.(*Hero).Name {
		return
	}
	if h.Party == e.One.(*Hero).Party {
		fmt.Println(h.GetHero(), "拍手叫好")
	}
	if h.Party != e.One.(*Hero).Party && h.Party != e.Another.(*Hero).Party {
		fmt.Println(h.GetHero(), "吃瓜群众")
	}
	if h.Party == e.Another.(*Hero).Party {
		fmt.Println(h.GetHero(), "过去帮助队友")
		// h.Fight(e.One, e.Not)
	}
}
func (h *Hero) Fight(l Listener, n Notifier) {
	e := &Event{
		Not:     n,
		One:     h,
		Another: l,
		Msg:     fmt.Sprintf("[%s]打了[%s]", h.GetHero(), l.(*Hero).GetHero()),
	}
	n.Notify(e)
}

type BaiXiaoSheng struct {
	heroList []Listener
}

func (b *BaiXiaoSheng) AddListener(l Listener) {
	b.heroList = append(b.heroList, l)
}
func (b *BaiXiaoSheng) RemoveListener(l Listener) {
	for i, v := range b.heroList {
		if v == l {
			b.heroList = append(b.heroList[:i], b.heroList[i+1:]...)
		}
	}
}
func (b *BaiXiaoSheng) Notify(e *Event) {
	fmt.Println("白晓生通知了：", e.Msg)
	for _, v := range b.heroList {
		v.OnFriendBeFight(e)
	}
}

func main() {
	b := &BaiXiaoSheng{}
	h1 := &Hero{
		Name:  "盖伦",
		Party: "德玛西亚",
	}
	h2 := &Hero{
		Name:  "赵信",
		Party: "德玛西亚",
	}
	h3 := &Hero{
		Name:  "嘉文四世",
		Party: "德玛西亚",
	}
	h4 := &Hero{
		Name:  "蛮王",
		Party: "诺克萨斯",
	}
	h5 := &Hero{
		Name:  "德莱文",
		Party: "诺克萨斯",
	}
	h6 := &Hero{
		Name:  "剑圣",
		Party: "艾欧尼亚",
	}
	b.AddListener(h1)
	b.AddListener(h2)
	b.AddListener(h3)
	b.AddListener(h4)
	b.AddListener(h5)
	b.AddListener(h6)
	h1.Fight(h4, b)
}
