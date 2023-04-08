package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/4/8
  @desc: 结构体
**/

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	var root treeNode
	fmt.Println(root)
}
