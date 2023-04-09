package main

import (
	"fmt"
	"learngo/object/tree"
)

/**
  @author: CodeWater
  @since: 2023/4/9
  @desc: 封装
		1. 一个目录下只能有一个包package
		2. 名字一般使用CamelCase
		3. 首字母大写表示：public
		4. 首字母小写表示：private
**/

func main() {
	var root tree.TreeNode
	fmt.Println(root)

	fmt.Println("==============================some init ways============================")

	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)

	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	fmt.Println("=============================createNode func=============================")
	root.Left.Right = tree.CreateNode(3)
	fmt.Println(root)

	fmt.Println("=============================treeNode's func=============================")
	root.Print()

	fmt.Println("=============================set value for treeNode=============================")
	root.Right.Left.SetValue(4)
	root.Right.Left.Print() //值不变
	root.SetValuePointer(100)
	root.Print() //值改变

	fmt.Println("=============================nil pointer call func=============================")
	node := tree.TreeNode{}
	node.Right.SetValuePointer(1) //Right空指针调用
	node.Print()

	fmt.Println("=============================traverse=============================")
	root.Traverse()
}
