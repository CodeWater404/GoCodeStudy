package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/8
  @desc: 适配器模式：
	另外一个例子：盖伦使用技能放出攻击；但是适配之后，变成盖伦放出攻击后电脑就关机。
**/

//适配目标，抽象的技能
type Attack interface {
	Fight()
}

//具体的技能
type Dabaojian struct {
}

func (d *Dabaojian) Fight() {
	fmt.Println("使用了大保健技能，将敌人击杀。。。。")
}

type Hero struct {
	Name   string
	attack Attack
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用了技能。。。")
	//使用具体的战斗方式
	h.attack.Fight()
}

//适配者
type PowerOff struct{}

func (p *PowerOff) ShutDown() {
	fmt.Println("电脑即将关机。。。。")
}

//适配器
type Adapter2 struct {
	poweroff *PowerOff
}

func (a *Adapter2) Fight() {
	a.poweroff.ShutDown()
}

func NewAdapter2(p *PowerOff) *Adapter2 {
	return &Adapter2{p}
}

func main() {
	//正常情况下，盖伦直接放技能攻击
	gailun := Hero{"盖伦", new(Dabaojian)}
	gailun.Skill()

	fmt.Println("=============================增加适配器=============================")
	//使用适配器之后：盖伦放技能攻击，然后电脑还会关机（adapter实现了attack接口所以可以传入，改变原有行为其实是adapter2内含另外一个
	//属性，然后这个属性再去重新实现attack接口里面的fight方法。因为盖伦skill里面调用了attack接口的抽象方法，所以具体运行的时候是换成新的）
	gailun = Hero{"盖伦2", NewAdapter2(new(PowerOff))}
	gailun.Skill()
}
