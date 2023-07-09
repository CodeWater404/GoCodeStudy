package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/9
  @desc: 模板方法模式

	例子：煮咖啡和泡茶都是同样的流程：煮开水、冲泡、倒入杯中、添加相应的料
**/

//抽象类，制作饮料，包裹一个模板，全部实现的步骤
type Beverage interface {
	BoilWater() //煮水
	PourInCup() //倒入杯中
	Brew()      //冲泡
	AddThings() //加料
	//新加一个方法来控制某个流程是否需要执行
	WantAddThings() bool
}

//封装一套流程模板基类，让具体的制作流程继承且实现
type template struct {
	b Beverage
}

//封装固定的模板，制作饮料
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	//固定的流程
	t.b.BoilWater() //这里真实运行的时候，会指向子类具体的实现方法
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

//具体的制作流程，制作咖啡（MakeCoffee继承template实现beverage接口）
type MakeCoffee struct {
	template //继承模板
}

//煮开水
func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

//冲泡
func (mc *MakeCoffee) Brew() {
	fmt.Println("用水冲咖啡豆")
}

//倒入杯中
func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将充好的咖啡倒入陶瓷杯中")
}

//添加辅料
func (mc *MakeCoffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

//咖啡需要添加辅料
func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	//b是Beverage，是makecoffee的接口，需要给接口赋值，让b指向具体的子类来触发b全部方法的多态特性
	makeCoffee.b = makeCoffee
	return makeCoffee
}

//同样的逻辑来制作茶
type MakeTea struct {
	template
}

//煮开水
func (mt *MakeTea) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

//冲泡
func (mt *MakeTea) Brew() {
	fmt.Println("用水冲茶叶")
}

//倒入杯中
func (mt *MakeTea) PourInCup() {
	fmt.Println("将充好的茶水倒入茶壶中")
}

//添加辅料
func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

//茶不需要添加辅料
func (mt *MakeTea) WantAddThings() bool {
	return false
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func main() {
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBeverage()

	fmt.Println("=============================制作茶=============================")
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}
