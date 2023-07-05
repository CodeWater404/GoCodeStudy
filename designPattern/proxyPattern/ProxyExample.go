package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/5
  @desc: 代理模式的另外一个例子
	潘金莲通过王婆的介绍，和西门大官人认识
**/

//抽象主题：美女都有的特性
type BeautyMoman interface {
	//抛媚眼
	MakeEyesWithMan()
	//美好时光
	HappyWithMan()
}

//具体的主题：潘金莲
type PanJinlian struct{}

func (p *PanJinlian) MakeEyesWithMan() {
	fmt.Println("潘金莲对本官抛了个媚眼")
}

func (p *PanJinlian) HappyWithMan() {
	fmt.Println("潘金莲和本官共度了美好的时光。。。。")
}

//中间代理人：王婆
type WangPo struct {
	woman BeautyMoman // 作为代理人是具有抽象的资源从而代理一大类型的事物，而不是具体的
}

//new实例的时候也是传入抽象的资源
func NewProxy2(woman BeautyMoman) BeautyMoman {
	return &WangPo{woman}
}

//代理人实现被代理人的功能
func (w *WangPo) MakeEyesWithMan() {
	w.woman.MakeEyesWithMan()
}

func (w *WangPo) HappyWithMan() {
	w.woman.HappyWithMan()
}

//具体的业务逻辑:王婆介绍潘金莲给大官人
func main() {
	wangPo := NewProxy2(new(PanJinlian))
	//实现被代理人的功能：王婆让潘金莲和大官人。。。。
	wangPo.HappyWithMan()
	wangPo.MakeEyesWithMan()
}
