package _1_100

/**
  @author: CodeWater
  @since: 2024/1/13
  @desc: 86. 分隔链表
**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	//很巧妙，开两个链表，左边存比x小的，右边存比x大的，最后把左边尾部指向右边的开头
	lh, rh := &ListNode{Val: -1}, &ListNode{Val: -1}
	lt, rt := lh, rh
	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			lt.Next = p
			lt = p
		} else {
			rt.Next = p
			rt = p
		}
	}
	lt.Next = rh.Next
	rt.Next = nil

	return lh.Next
}
