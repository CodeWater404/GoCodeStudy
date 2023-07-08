package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/8
  @desc: 外观模式的另外一个例子

	家庭里面同样还是会有许多的电器，通过一个共有的家庭影院去打开不同的模式，唤醒和关闭不同的设备
**/

//电视机
type TV struct{}

func (t *TV) On() {
	fmt.Println("open TV")
}

func (t *TV) Off() {
	fmt.Println("close TV")
}

//音箱
type VoiceBox struct{}

func (vb *VoiceBox) On() {
	fmt.Println("open VoiceBox")
}

func (vb *VoiceBox) Off() {
	fmt.Println("close VoiceBox")
}

//灯光
type Light struct{}

func (l *Light) On() {
	fmt.Println("open Light")
}

func (l *Light) Off() {
	fmt.Println("close Light")
}

//游戏机
type Xbox struct{}

func (x *Xbox) On() {
	fmt.Println("open xbox")
}

func (x *Xbox) Close() {
	fmt.Println("close xbox")
}

//麦克风
type MicroPhone struct{}

func (mp *MicroPhone) On() {
	fmt.Println("open MicroPhone")
}

func (mp *MicroPhone) Close() {
	fmt.Println("close MicroPhone")
}

//投影仪
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("open Projector")
}

func (p *Projector) Close() {
	fmt.Println("close Projector")
}

//家庭影院（外观）
type HomePlayerFacode struct {
	tv    TV
	vb    VoiceBox
	light Light
	xbox  Xbox
	mp    MicroPhone
	pro   Projector
}

//ktv模式
func (hp *HomePlayerFacode) DoKTV() {
	fmt.Println("homePlayerFacode enter KTV mode")
	hp.tv.On()
	hp.pro.On()
	hp.mp.On()
	hp.light.Off()
	hp.vb.On()
}

//游戏模式
func (hp *HomePlayerFacode) DoGame() {
	fmt.Println("homePlayerFacode enter Game mode")
	hp.tv.On()
	hp.light.On()
	hp.xbox.On()
}

func main() {
	homePlayer := new(HomePlayerFacode)
	homePlayer.DoKTV()

	fmt.Println("=============================game mode=============================")
	homePlayer.DoGame()
}
