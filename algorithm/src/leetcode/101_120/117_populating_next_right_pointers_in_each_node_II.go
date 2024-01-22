package _01_120

/**
  @author: CodeWater
  @since: 2024/1/22
  @desc: 117. 填充每个节点的下一个右侧节点指针 II
**/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	cur := root
	for cur != nil {
		//用一个单链表维护下一层的所有节点，在这个处理过程中实际上就是把一层的next赋完值了
		head := &Node{Val: -1}
		tail := head
		for p := cur; p != nil; p = p.Next {
			if p.Left != nil {
				//左节点不空，加入到链表中，然后尾指针往后移动
				tail.Next = p.Left
				tail = tail.Next
			}
			if p.Right != nil {
				tail.Next = p.Right
				tail = tail.Next
			}

		}
		//往下一层移动,因为单链表的第一个节点是保存的下一层的节点信息
		cur = head.Next
	}
	return root
}
