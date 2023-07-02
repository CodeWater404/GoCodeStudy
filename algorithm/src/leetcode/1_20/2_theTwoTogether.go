package __20

/**
  @author: CodeWater
  @since: 2023/7/2
  @desc: 两数相加
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//head := new(ListNode)
	head := &ListNode{}
	current := head
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		//current.Next = new(ListNode)
		current.Next = &ListNode{Val: carry % 10}
		current = current.Next
		carry = carry / 10
	}
	return head.Next

}