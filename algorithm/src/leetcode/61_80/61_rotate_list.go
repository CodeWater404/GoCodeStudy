package _1_80

/**
  @author: CodeWater
  @since: 2024/1/13
  @desc: 61. 旋转链表
**/

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n, tail := 0, &ListNode{} //n链表长度，tail链表尾部
	for p := head; p != nil; p = p.Next {
		tail = p
		n++
	}
	k = k % n
	if k == 0 {
		return head
	}
	p := head
	for i := 0; i < n-k-1; i++ {
		p = p.Next
	}
	tail.Next = head //尾部指向开头
	head = p.Next    //更新头
	p.Next = nil     //新的尾部
	return head

}
