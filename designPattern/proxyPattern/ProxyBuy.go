package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/5
  @desc: 代理模式
	优点：
		(1)能够协调调用者和被调用者，在一定程度上降低了系统的耦合度。
		(2)客户端可以针对抽象主题角色进行编程，增加和更换代理类无须修改源代码，符合开闭原则，系统具有较好的灵活性和可扩展性。

	缺点：
		(1)代理实现较为复杂。
**/

//商品： 种类  真假
type Goods struct {
	Kind string
	Fact bool
}

//==============抽象层==================
type Shopping interface {
	Buy(goods *Goods)
}

//==============实现层==================
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国购买了：", goods.Kind)
}

type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国购买了：", goods.Kind)
}

type AfricanShopping struct{}

func (as *AfricanShopping) Buy(goods *Goods) {
	fmt.Println("去非洲购买了：", goods.Kind)
}

//海外代理
type OverSeaProxy struct {
	shopping Shopping // 代理某一类主题，这里是抽象类型
}

//ps： 在Go语言中，接口类型的变量可以持有任意实现了该接口的具体类型的值。这意味着，
//当一个具体类型实现了接口中定义的全部方法时，该具体类型的指针也可以被赋值给对应的接口变量。
func NewProxy(shopping Shopping) Shopping {
	return &OverSeaProxy{shopping}
}

func (op *OverSeaProxy) Buy(goods *Goods) {
	if op.distinguish(goods) == true {
		op.shopping.Buy(goods) //多态
		op.check(goods)
	}
}

//辨别真假的能力
func (op *OverSeaProxy) distinguish(goods *Goods) bool {
	fmt.Println("=========对【", goods.Kind, "】==============进行了辨别")
	if goods.Fact == false {
		fmt.Println("!!!!发现了假货：", goods.Kind, ", 不值得购买")
	}
	return goods.Fact
}

//海关安检流程
func (op *OverSeaProxy) check(goods *Goods) {
	fmt.Println("对【", goods.Kind, "】进行了海关安检，成功购买！！！")
}

//==============业务逻辑层=========================
func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}
	g2 := Goods{
		Kind: "cet证书",
		Fact: false,
	}

	var kShopping Shopping
	kShopping = new(KoreaShopping)
	var k_proxy Shopping
	k_proxy = NewProxy(kShopping)
	k_proxy.Buy(&g1)

	var aShopping Shopping
	aShopping = new(AmericanShopping)
	/*原视频中是重新new了一个代理：
	var a_proxy Shopping
	a_proxy = NewProxy(aShopping)
	a_proxy.Buy(&g2)
	*/
	k_proxy = NewProxy(aShopping)
	k_proxy.Buy(&g2)
	/*对于原教程中new了两个不同的shopping：aShopping和kShopping，我猜测是因为区分有不同的客户需要购买（因为不用
	声明两个变量直接重新new一个也是可以的）
	代理new两个好像也是同理，表现有不同的代理人。
	但我的想法是，抽象层都是代理人，实现再new成不同类型的代理人，所以不需
	要声明两个变量。
	不知道我的想法正不正确。
	*/
}
