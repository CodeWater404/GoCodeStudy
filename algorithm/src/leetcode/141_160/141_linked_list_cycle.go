package _41_160

/**
  @author: CodeWater
  @since: 2024/1/8
  @desc: 141. 环形链表
**/

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil {
		slow, fast = slow.Next, fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if slow == fast {
			return true
		}
	}
	return false
}
