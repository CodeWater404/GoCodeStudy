package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/8
  @desc: 外观模式
	优缺点：
		优点：
			(1)它对客户端屏蔽了子系统组件，减少了客户端所需处理的对象数目，并使得子系统使用
			起来更加容易。通过引入外观模式，客户端代码将变得很简单，与之关联的对象也很少。
			(2)它实现了子系统与客户端之间的松耦合关系，这使得子系统的变化不会影响到调用它的
			客户端，只需要调整外观类即可。
			(3)一个子系统的修改对其他子系统没有任何影响。
		缺点：
			(1)不能很好地限制客户端直接使用子系统类，如果对客户端访问子系统类做太多的限制则
			减少了可变性和灵活性。
			(2)如果设计不当，增加新的子系统可能需要修改外观类的源代码，违背了开闭原则。

	适用场景：
		(1)复杂系统需要简单入口使用。
		(2)客户端程序与多个子系统之间存在很大的依赖性。
		(3)在层次化结构中，可以使用外观模式定义系统中每一层的入口，层与层之间不直接产生
		联系，而通过外观类建立联系，降低层之间的耦合度。

	例子： 洗衣机、电视、扫地机器人各自有不同的遥控器，通过给他们提供一个公共的平台让他们在这个上面实现各自的功能，然后通过平台的遥控器去
		 操控平台，平台在调用对应的系统（电器）。
**/

type SubSystemA struct{}

func (sa *SubSystemA) MethodA() {
	fmt.Println("subsystem method A")
}

type SubSystemB struct{}

func (sb SubSystemB) MethodB() {
	fmt.Println("subsystem method B")
}

type SubSystemC struct{}

func (sc *SubSystemC) MethodC() {
	fmt.Println("subsystem method C")
}

type SubSystemD struct{}

func (sd *SubSystemD) MethodD() {
	fmt.Println("subsystem method D")
}

//外观模式： 提供一个外观类，简化成一个简单的接口提供使用
type Facode struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

func (f *Facode) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f *Facode) MethodTwo() {
	f.c.MethodC()
	f.d.MethodD()
}

func main() {
	fmt.Println("=============================not use appearance pattern to call method a and b=========================")
	sa := new(SubSystemA)
	sa.MethodA()
	sb := new(SubSystemB)
	sb.MethodB()

	fmt.Println("=============================use appearance pattern to call method a and b=============================")
	f := Facode{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
		d: new(SubSystemD),
	}
	//调用外观模式的包裹方法实现
	f.MethodOne()
}
