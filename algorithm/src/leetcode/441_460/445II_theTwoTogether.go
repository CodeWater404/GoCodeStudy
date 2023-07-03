package _41_460

/**
  @author: CodeWater
  @since: 2023/7/3
  @desc: 两数相加II
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)
	head := &ListNode{Val: -1}
	t := 0
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		cur := &ListNode{Val: t % 10}
		t /= 10
		//头插法，最后少一次翻转
		cur.Next = head.Next
		head.Next = cur
	}

	return head.Next
}

func reverse(l *ListNode) *ListNode {
	a, b := l, l.Next
	for b != nil {
		c := b.Next
		b.Next = a
		a, b = b, c
	}
	l.Next = nil
	return a
}