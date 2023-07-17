package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/17
  @desc: 抽象工厂模式：
	产品族：具有同一个地区、同一个厂商、同一个开发包、同一个组织模块等，但是具备不同特点或功能的产品集合，称之为是一个产品族。
	产品等级结构：具有相同特点或功能，但是来自不同的地区、不同的厂商、不同的开发包、不同的组织模块等的产品集合，称之为是一个产品等级结构。
	当程序中的对象可以被划分为产品族和产品等级结构之后，那么“抽象工厂方法模式”才可以被适用。
	优点：
		1.  拥有工厂方法模式的优点
		2. 当一个产品族中的多个对象被设计成一起工作时，它能够保证客户端始终只使用同一个产品族中的对象。
		3   增加新的产品族很方便，无须修改已有系统，符合“开闭原则”。
	缺点：
		1. 增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，甚至需要修改抽象层代码，这显然会带来较大的不便，违背了“开闭原则”。
	适用场景：
		(1) 系统中有多于一个的产品族。而每次只使用其中某一产品族。可以通过配置文件等方式来使得用户可以动态改变产品族，也可以很方便地增加新的产品族。
		(2) 产品等级结构稳定。设计完成之后，不会向系统中增加新的产品等级结构或者删除已有的产品等级结构。


	例子：还是工厂生产水果的。但是水果种类相对固定了，形成产品族、产品等级
**/

//抽象层
type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

//抽象工厂
type AbstractFactory2 interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

//中国产品族
type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("中国苹果")
}

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("中国香蕉")
}

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() {
	fmt.Println("中国梨")
}

type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

//日本产品族
type JapanApple struct{}

func (ja *JapanApple) ShowApple() {
	fmt.Println("日本苹果")
}

type JapanBanana struct{}

func (jb *JapanBanana) ShowBanana() {
	fmt.Println("日本香蕉")
}

type JapanPear struct{}

func (jp *JapanPear) ShowPear() {
	fmt.Println("日本梨")
}

type JapanFactory struct{}

func (jf *JapanFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(JapanApple)
	return apple
}

func (jf *JapanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(JapanBanana)
	return banana
}

func (jf *JapanFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(JapanPear)
	return pear

}

//具体的业务逻辑层
func main() {
	//业务1：如果只是需要中国的苹果、香蕉、梨，那么只需要new中国的工厂即可
	var cf AbstractFactory2
	cf = new(ChinaFactory)
	//中国苹果
	var cApple AbstractApple
	cApple = cf.CreateApple()
	cApple.ShowApple()
	//中国香蕉
	var cBanana AbstractBanana
	cBanana = cf.CreateBanana()
	cBanana.ShowBanana()
	//中国梨
	var cPear AbstractPear
	cPear = cf.CreatePear()
	cPear.ShowPear()

	fmt.Println("==========================================================")
	//业务2：中国梨和日本香蕉，这个时候就需要两个工厂
	var jf AbstractFactory2
	jf = new(JapanFactory)
	var jBanana AbstractBanana
	jBanana = jf.CreateBanana()
	jBanana.ShowBanana()

	//这里为了方便就直接用上面的代码调一下
	cPear.ShowPear()
}
