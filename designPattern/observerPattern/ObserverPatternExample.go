package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/14
  @desc:
	例子：
		假设江湖有一名无事不知，无话不说的大嘴巴，“江湖百晓生”，任何江湖中发生的事件都会被百晓生知晓，且进行广播。
		先江湖中有两个帮派，分别为：
		丐帮：黄蓉、洪七公、乔峰。
		明教：张无忌、灭绝师太、金毛狮王。
		现在需要用观察者模式模拟如下场景：
		（1）事件一：丐帮的黄蓉把明教的张无忌揍了，这次武林事件被百晓生知晓，并且进行广播。
				 主动打人方的帮派收到消息要拍手叫好。
				 被打的帮派收到消息应该报酬，如：灭绝师太得知消息进行报仇，将丐帮黄蓉揍了。触发事件二。
		（2）事件二：明教的灭绝师太把丐帮的黄蓉揍了，这次武林事件被百姓生知晓，并且进行广播。
	...
**/

const (
	PGaiBang  string = "丐帮"
	PMingJiao string = "明教"
)

//抽象层
type Listener2 interface {
	//当同伴被揍了怎么办
	OnFriendBeFight(event *Event)
	GetName() string
	GetParty() string
	Title() string
}

type Notifier2 interface {
	//添加观察者
	AddListener(listener2 Listener2)
	//删除观察者
	RemoveListener(listener2 Listener2)
	//通知广播
	Notify(event *Event)
}

type Event struct {
	Noti    Notifier2 //被知晓的通知者
	One     Listener2 //事件主动发出者
	Another Listener2 //事件被动接收者
	Msg     string    //具体的消息

}

//实现层
//英雄（Listener）
type Hero struct {
	Name  string
	Party string
}

func (h *Hero) Fight(another Listener2, baixiao Notifier2) {
	msg := fmt.Sprintf("%s 将 %s 揍了。。。", h.Title(), another.Title())

	//生成事件
	event := new(Event)
	event.Noti = baixiao
	event.One = h
	event.Another = another
	event.Msg = msg

	baixiao.Notify(event)
}

func (h *Hero) Title() string {
	return fmt.Sprintf("[%s]:%s", h.Party, h.Name)
}

func (h *Hero) OnFriendBeFight(event *Event) {
	//判断是否为当事人
	if h.Name == event.One.GetName() || h.Name == event.Another.GetName() {
		return
	}

	//同一帮派的人把别人揍了，拍手叫好
	if h.Party == event.One.GetParty() {
		fmt.Println(h.Title(), "得知消息，拍手叫好！！！！")
		return
	}

	//被别的帮派揍了，主动报仇反击
	if h.Party == event.Another.GetParty() {
		fmt.Println(h.Title(), "得知消息，发起复仇反击！！！！！")
		//这里一旦开启就会死循环，两帮人马打来打去
		h.Fight(event.One, event.Noti)
		return
	}
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) GetParty() string {
	return h.Party
}

//百晓生（notifier）
type BaiXiao struct {
	heroList []Listener2
}

//添加观察者
func (b *BaiXiao) AddListener(listener Listener2) {
	b.heroList = append(b.heroList, listener)
}

//删除观察者
func (b *BaiXiao) RemoveListener(listener2 Listener2) {
	for index, l := range b.heroList {
		if listener2 == l {
			b.heroList = append(b.heroList[:index], b.heroList[index+1:]...)
			break
		}
	}
}

//通知广播
func (b *BaiXiao) Notify(event *Event) {
	fmt.Println("【世界消息】 百晓生广播消息： ", event.Msg)
	for _, listener := range b.heroList {
		//依次调用全部的观察者的具体动作
		listener.OnFriendBeFight(event)
	}
}

//业务逻辑层
func main() {
	hero1 := Hero{"黄容", PGaiBang}
	hero2 := Hero{"洪七公", PGaiBang}
	hero3 := Hero{"乔峰", PGaiBang}
	hero4 := Hero{"张无忌", PMingJiao}
	hero5 := Hero{"韦一笑", PMingJiao}
	hero6 := Hero{"金毛狮王", PMingJiao}

	baixiao := BaiXiao{}
	baixiao.AddListener(&hero1)
	baixiao.AddListener(&hero2)
	baixiao.AddListener(&hero3)
	baixiao.AddListener(&hero4)
	baixiao.AddListener(&hero5)
	baixiao.AddListener(&hero6)

	fmt.Println("武林一片平静。。。。。")
	hero1.Fight(&hero4, &baixiao)
}
