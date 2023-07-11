package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/11
  @desc: 命令模式例子二
	路边撸串烧烤场景， 有烤羊肉，烧鸡翅命令，有烤串师傅，和服务员MM。根据命令模式，设计烤串场景。

**/

//烤串师傅
type Cooker struct{}

func (c *Cooker) MakeChicken() {
	fmt.Println("烤串似乎烤了鸡肉串")
}

func (c *Cooker) MakeChuaner() {
	fmt.Println("烤串师傅烤了羊肉串")
}

//抽象的命令
type Command interface {
	Make()
}

//鸡肉串订单
type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeChicken()
}

//羊肉串订单
type CommandCookChuaner struct {
	cooker *Cooker
}

func (cmd *CommandCookChuaner) Make() {
	cmd.cooker.MakeChuaner()
}

//服务员：收集订单命令,然后通知给烤串师傅
type WaiterMM struct {
	CmdList []Command //命令的集合
}

func (w *WaiterMM) Notify() {
	if w.CmdList == nil {
		return
	}

	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}

func main() {
	cooker := new(Cooker)
	cmdChicken := CommandCookChicken{cooker}
	cmdChuaner := CommandCookChuaner{cooker}

	waiter := new(WaiterMM)
	waiter.CmdList = append(waiter.CmdList, &cmdChicken, &cmdChuaner)
	waiter.Notify()
}
