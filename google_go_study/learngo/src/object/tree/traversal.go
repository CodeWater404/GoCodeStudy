package tree

/**
  @author: CodeWater
  @since: 2023/4/9
  @desc: 演示把一个结构体的方法放在不同文件里面，但是都是在同一个包下
**/

/**traverse
** @Description: 中序遍历
** @receiver node
**/
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse() //这里如果是别的语言还要判断Left是否为空，但是go不需要，因为nil也可调用函数
	node.Print()
	node.Right.Traverse()

}
