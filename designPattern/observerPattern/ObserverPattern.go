package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/13
  @desc: 观察者模式
	优点：
		(1) 观察者模式可以实现表示层和数据逻辑层的分离，定义了稳定的消息更新传递机制，并抽象了更新接口，使得可以有各种各样不同的表示层充当具体观察者角色。
		(2) 观察者模式在观察目标和观察者之间建立一个抽象的耦合。观察目标只需要维持一个抽象观察者的集合，无须了解其具体观察者。由于观察目标和观察者没有紧密地耦合在一起，因此它们可以属于不同的抽象化层次。
		(3) 观察者模式支持广播通信，观察目标会向所有已注册的观察者对象发送通知，简化了一对多系统设计的难度。
		(4) 观察者模式满足“开闭原则”的要求，增加新的具体观察者无须修改原有系统代码，在具体观察者与观察目标之间不存在关联关系的情况下，增加新的观察目标也很方便。
	缺点：
		(1) 如果一个观察目标对象有很多直接和间接观察者，将所有的观察者都通知到会花费很多时间。
		(2) 如果在观察者和观察目标之间存在循环依赖，观察目标会触发它们之间进行循环调用，可能导致系统崩溃。
		(3) 观察者模式没有相应的机制让观察者知道所观察的目标对象是怎么发生变化的，而仅仅只是知道观察目标发生了变化。

	5.4.5 适用场景
		(1) 一个抽象模型有两个方面，其中一个方面依赖于另一个方面，将这两个方面封装在独立的对象中使它们可以各自独立地改变和复用。
		(2) 一个对象的改变将导致一个或多个其他对象也发生改变，而并不知道具体有多少对象将发生改变，也不知道这些对象是谁。
		(3) 需要在系统中创建一个触发链，A对象的行为将影响B对象，B对象的行为将影响C对象……，可以使用观察者模式创建一种链式触发机制。

	例子： 学生在干自己的事，班长观察老师是否来了；老师一来，班长通知学生，大家停止自己的事。
**/

//抽象层
type Listener interface {
	OnTeacherComming() //观察者得到通知后要触发的动作
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

//实现层
//观察者，学生
type Stu1 struct {
	BadThing string
}

func (s *Stu1) OnTeacherComming() {
	fmt.Println("stu1 停止", s.BadThing)
}

func (s *Stu1) DoBadthing() {
	fmt.Println("stu1 正在", s.BadThing)
}

type Stu2 struct {
	BadThing string
}

func (s *Stu2) OnTeacherComming() {
	fmt.Println("stu2 停止", s.BadThing)
}

func (s *Stu2) DoBadthing() {
	fmt.Println("stu2 正在", s.BadThing)
}

type Stu3 struct {
	BadThing string
}

func (s *Stu3) OnTeacherComming() {
	fmt.Println("stu3 停止", s.BadThing)
}

func (s *Stu3) DoBadthing() {
	fmt.Println("stu3 正在", s.BadThing)
}

//通知者班长
type ClassMonitor struct {
	listenerList []Listener //需要通知的全部观察者的集合
}

func (m *ClassMonitor) AddListener(listener Listener) {
	m.listenerList = append(m.listenerList, listener)
}

func (m *ClassMonitor) RemoveListener(listener Listener) {
	for index, l := range m.listenerList {
		if listener == l {
			m.listenerList = append(m.listenerList[:index], m.listenerList[index+1:]...)
			break
		}
	}
}

func (m *ClassMonitor) Notify() {
	for _, listener := range m.listenerList {
		//依次调用全部观察的具体动作
		listener.OnTeacherComming()
	}
}

func main() {
	s1 := &Stu1{BadThing: "抄作业"}
	s2 := &Stu2{BadThing: "玩王者荣耀"}
	s3 := &Stu3{BadThing: "看别人玩王者荣耀"}

	classMonitor := new(ClassMonitor)
	fmt.Println("上课了，但是老师没有来，学生们都在忙自己的事情。。。。")
	s1.DoBadthing()
	s2.DoBadthing()
	s3.DoBadthing()

	classMonitor.AddListener(s1)
	classMonitor.AddListener(s2)
	classMonitor.AddListener(s3)

	fmt.Println("老师来了，班长给学生们使了一个眼色")
	classMonitor.Notify()

}
