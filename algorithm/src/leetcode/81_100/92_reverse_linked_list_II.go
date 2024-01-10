package _1_100

/**
  @author: CodeWater
  @since: 2024/1/9
  @desc: 92. 反转链表 II
**/

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	a := dummy
	for i := 0; i < left-1; i++ { //走到left位置前一个节点
		a = a.Next
	}
	b := a.Next
	c := b.Next
	//left到right中间的交换链表指向顺序
	for i := 0; i < right-left; i++ {
		d := c.Next
		c.Next, b, c = b, c, d
	}
	//反转left和right位置：left位置的节点指向right的下一个节点c
	a.Next.Next = c
	//a节点指向right位置的节点
	a.Next = b
	return dummy.Next
}
