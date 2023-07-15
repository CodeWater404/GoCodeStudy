package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/15
  @desc: 简单工厂模式
	优点：
		1. 实现了对象创建和使用的分离。
		2. 不需要记住具体类名，记住参数即可，减少使用者记忆量。
	缺点：
		1. 对工厂类职责过重，一旦不能工作，系统受到影响。
		2. 增加系统中类的个数，复杂度和理解度增加。
		3. 违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂。
	适用场景：
		1.  工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
		2. 客户端只知道传入工厂类的参数，对于如何创建对象并不关心。

	例子：工厂会生产出不同的水果，业务层只需要跟工厂交互，不需要自己去创建具体的水果
**/

//抽象层
//水果类（抽象接口）
type Fruit interface {
	Show()
}

//基础类模块
type Apple struct {
	Fruit //这行其实可以省略，只要实现接口即可，为了易于理解显示继承
}

func (a *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct {
	Fruit
}

func (b *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {
	Fruit
}

func (p *Pear) Show() {
	fmt.Println("我是梨")
}

//工厂模块
type Factory struct {
}

//一个工厂有一个生产水果的机器，返回一个抽象水果的指针，但是实际指向的是具体的水果
func (f *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "苹果" {
		fruit = new(Apple)
	} else if kind == "香蕉" {
		fruit = new(Banana)
	} else if kind == "梨" {
		fruit = new(Pear)
	}

	return fruit
}

//具体的业务层
func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("苹果")
	apple.Show()

	banana := factory.CreateFruit("香蕉")
	banana.Show()

	pear := factory.CreateFruit("梨")
	pear.Show()
}
