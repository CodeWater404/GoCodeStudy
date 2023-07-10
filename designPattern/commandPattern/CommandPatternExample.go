package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/10
  @desc: $

	进阶例子：病人不通过病单联系医生，而是通过护士，然后护士查看病单再联系医生
**/

//医生，命令的接收者
type Doctor2 struct{}

func (d *Doctor2) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor2) treatNose() {
	fmt.Println("医生治疗鼻子")
}

//抽象的命令
type Command2 interface {
	Treat2()
}

//治疗眼睛的病单
type CommandTreatEye2 struct {
	doctor *Doctor2
}

func (cmd *CommandTreatEye2) Treat2() {
	cmd.doctor.treatEye()
}

//治疗鼻子的病单
type CommandTreatNose2 struct {
	doctor *Doctor2
}

func (cmd *CommandTreatNose2) Treat2() {
	cmd.doctor.treatNose()
}

//护士---调用命令者
type Nurse struct {
	CmdList []Command2 //收集的命令集合
}

//发送病单，发送命令的方法
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		//执行病单绑定的命令（这里会调用病单已经绑定的医生的诊断的方法）
		cmd.Treat2()
	}
}

//病人——业务层
func main() {
	//依赖病单，通过填写病单，让医生看病
	doctor := new(Doctor2)
	//治疗眼睛的病单
	cmdEye := CommandTreatEye2{doctor}
	//治疗鼻子的病单
	cmdNose := CommandTreatNose2{doctor}

	//护士
	nurse := new(Nurse)
	//收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye, &cmdNose)

	//执行收集病单的指令
	nurse.Notify()
}
