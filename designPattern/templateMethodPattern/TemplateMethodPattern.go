package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/9
  @desc: 模板方法模式
	优缺点：
		优点：
			(1)在父类中形式化地定义一个算法，而由它的子类来实现细节的处理，在子类实现详细的
			处理算法时并不会改变算法中步骤的执行次序。
			(2)模板方法模式是一种代码复用技术，它在类库设计中尤为重要，它提取了类库中的公共
			行为，将公共行为放在父类中，而通过其子类来实现不同的行为，它鼓励我们恰当使用继
			承来实现代码复用。
			(3)可实现一种反向控制结构，通过子类覆盖父类的钩子方法来决定某一特定步骤是否需要
			执行。
			(4)在模板方法模式中可以通过子类来覆盖父类的基本方法，不同的子类可以提供基本方法
			的不同实现，更换和增加新的子类很方便，符合单一职责原则和开闭原则。
		缺点：
			需要为每一个基本方法的不同实现提供一个子类，如果父类中可变的基本方法太多，将会
			导致类的个数增加，系统更加庞大，设计也更加抽象。
	适用场景：
		(1)具有统一的操作步骤或操作过程；
		(2)具有不同的操作细节；
		(3)存在多个具有同样操作步骤的应用场景，但某些具体的操作细节却各不相同；
		在抽象类中统一操作步骤，并规定好接口；让子类实现接口。这样可以把各个具体的子类
		和操作步骤解耦合。

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
