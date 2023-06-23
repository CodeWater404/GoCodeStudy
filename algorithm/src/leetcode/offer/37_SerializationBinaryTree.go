package offer

import (
	"strconv"
	"strings"
)

/**
  @author: CodeWater
  @since: 2023/6/23
  @desc: 序列化二叉树
**/
//bfs
func serialize(root *TreeNode) string {
	ans := "["
	//中序遍历
	nodeQueue := make([]*TreeNode, 0)
	nodeQueue = append(nodeQueue, root)

	for len(nodeQueue) != 0 {
		tmpNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		if tmpNode == nil {
			ans += "null" + ","
			continue
		} else {
			ans += strconv.Itoa(tmpNode.Val) + ","
		}

		nodeQueue = append(nodeQueue, tmpNode.Left)
		nodeQueue = append(nodeQueue, tmpNode.Right)
	}
	//这里去掉最后一个元素应该是为了去掉逗号
	ans = ans[:len(ans)-1]
	return ans + "]"
}

func deserialize(data string) *TreeNode {
	//去除“【】”
	data = data[1 : len(data)-1]
	//分割元素，因为他的输入时字符串形式的数组
	dataSplit := strings.Split(data, ",")

	//空。但凡有一个结点都是有三个元素，val、null、null
	if len(dataSplit) == 1 {
		return nil
	}

	nodeQueue := make([]*TreeNode, 0)
	dataSplitCount := 0
	rootVal, _ := strconv.Atoi(dataSplit[dataSplitCount])
	root := &TreeNode{Val: rootVal, Left: nil, Right: nil}
	dataSplitCount++
	nodeQueue = append(nodeQueue, root)

	for len(nodeQueue) != 0 {
		tmpNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		leftNodeContext := dataSplit[dataSplitCount]
		dataSplitCount++
		if leftNodeContext == "null" {
			tmpNode.Left = nil
		} else {
			leftVal, _ := strconv.Atoi(leftNodeContext)
			leftNode := &TreeNode{Val: leftVal, Left: nil, Right: nil}
			tmpNode.Left = leftNode
			nodeQueue = append(nodeQueue, leftNode)
		}

		rightNodeContext := dataSplit[dataSplitCount]
		dataSplitCount++
		if rightNodeContext == "null" {
			tmpNode.Right = nil
		} else {
			rightVal, _ := strconv.Atoi(rightNodeContext)
			rightNode := &TreeNode{Val: rightVal, Left: nil, Right: nil}
			tmpNode.Right = rightNode
			nodeQueue = append(nodeQueue, rightNode)
		}
	}

	return root
}
