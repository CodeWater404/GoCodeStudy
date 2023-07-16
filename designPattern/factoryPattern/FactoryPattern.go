package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/16
  @desc: 工厂方法模式：在简单工厂模式基础上，把工厂再抽象一层，符合开闭原则！！！！
		也可以说： 简单工厂 + 开闭原则 = 工厂方法

	例子： 工厂会生产出不同的水果，业务层只需要跟工厂交互，不需要自己去创建具体的水果
**/

//抽象层
//水果类（抽象接口）
type Fruit2 interface {
	Show()
}

//工厂类，抽象接口
type AbstractFactory interface {
	CreateFruit() Fruit2 //生产水果类（抽象）的生产器方法
}

//基础类模块
type Apple2 struct {
	Fruit2 //为了易于理解，这里显示继承（其实可以省略）
}

func (a *Apple2) Show() {
	fmt.Println("我是苹果")
}

type Banana2 struct {
	Fruit2 //为了易于理解，这里显示继承（其实可以省略）
}

func (a *Banana2) Show() {
	fmt.Println("我是香蕉")
}

type Pear2 struct {
	Fruit2 //为了易于理解，这里显示继承（其实可以省略）
}

func (a *Pear2) Show() {
	fmt.Println("我是梨")
}

//=================新增一个水果种类
type JapanPear struct {
	Fruit2
}

func (jp *JapanPear) Show() {
	fmt.Println("我是日本的梨。。。。")
}

//工厂模块
//具体的苹果工厂
type AppleFactory struct {
	AbstractFactory //这里同样是为了理解，显示继承
}

func (af *AppleFactory) CreateFruit() Fruit2 {
	var fruit Fruit2
	//生产一个具体的苹果
	fruit = new(Apple2)
	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (bf *BananaFactory) CreateFruit() Fruit2 {
	var fruit Fruit2
	fruit = new(Banana2)
	return fruit
}

type PearFactory struct {
	AbstractFactory
}

func (pf *PearFactory) CreateFruit() Fruit2 {
	var fruit Fruit2
	fruit = new(Pear2)
	return fruit
}

//=============新增的一种水果
type JapanPearFactory struct {
	AbstractFactory
}

func (jpf *JapanPearFactory) CreateFruit() Fruit2 {
	var fruit Fruit2
	fruit = new(JapanPear)
	return fruit
}

//业务逻辑层
func main() {
	//本案例为了突出根据依赖倒转原则与面向接口编程特性. 一些变量的定义将使用显示类型声明方式
	//1. 生产一个具体的苹果对象：先要一个具体的苹果工厂；生产相对应的水果
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)
	//声明的fruit2接口，但是通过工厂产生真正的苹果对象
	var apple Fruit2
	apple = appleFac.CreateFruit()
	apple.Show()

	//2.生产一个香蕉和梨
	var banananFac AbstractFactory
	banananFac = new(BananaFactory)
	var banana Fruit2
	banana = banananFac.CreateFruit()
	banana.Show()
	//简洁写法，但是不易于理解其中的设计模式
	pearFac := new(PearFactory)
	pear := pearFac.CreateFruit()
	pear.Show()

	//如果现在相应新增加一个日本的梨，那么只需要整对应的具体水果、生产水果的抽象工厂
	var japanPearFac AbstractFactory
	japanPearFac = new(JapanPearFactory)
	var japanPear Fruit2
	japanPear = japanPearFac.CreateFruit()
	japanPear.Show()
}
