package tree

import "fmt"

/**
  @author: CodeWater
  @since: 2023/4/8
  @desc: 结构体
		1. go语言没有继承和多态、没有构造函数；一切都是采用结构体来实现
		2. nil指针也可以调用方法！！！！
		3. 要改变内容，必须使用指针接收者(写在函数名前的变量)
		4. 结构过大也可以考虑指针接收者
		5. 一致性：如有指针接收者，最好都是指针接收者
		6. 值/指针接收者均可接收值/指针,如：
			func (node *TreeNode) hello(){}
			node1 := TreeNode{}
			node2 := *TreeNode{}
			node1和node2都可以调用hello函数
**/

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

/**Print
** @Description: 类似于treeNode的一个函数
** @receiver node
**/
func (node TreeNode) Print() {
	fmt.Print(node.Value, "\n")
}

/**CreateNode
** @Description: 有时候结构体确实需要构造函数，那么这时候通过创建一个函数去实现.注意这里返回的是函数内部（也就是局部）变量的一个地址，但是对于go语言来说，不影响外部变量的使用。
** @param value
** @return *TreeNode
**/
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

/**setValue
** @Description: 接收者node是值传递，不会改变值
** @receiver node
** @param value
**/
func (node TreeNode) SetValue(value int) {
	node.Value = value
}

/**setValuePointer
** @Description: 接收者node是值传递，但是是指针类型，会改变传进来变量value的值
** @receiver node
** @param value
**/
func (node *TreeNode) SetValuePointer(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value
}
