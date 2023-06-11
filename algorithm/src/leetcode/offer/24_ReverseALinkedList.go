package offer

/**
  @author: CodeWater
  @since: 2023/6/9
  @desc: 反转链表
**/

// ListNode
//  @Description: Definition for singly-linked list.
//
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//声明空节点
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}
