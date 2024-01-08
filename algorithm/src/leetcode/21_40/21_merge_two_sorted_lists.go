package _1_40

/**
  @author: CodeWater
  @since: 2024/1/8
  @desc: 21. 合并两个有序链表
**/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		} else {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
