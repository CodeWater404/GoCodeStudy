package _1_100

/**
  @author: CodeWater
  @since: 2024/1/11
  @desc: 82. 删除排序链表中的重复元素 II
**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	p := dummy
	for p.Next != nil {
		q := p.Next.Next
		//扫描一段区间【p.Next,q)
		for q != nil && q.Val == p.Next.Val {
			q = q.Next
		}
		//p.Next到q这一段没有重复的数
		if p.Next.Next == q {
			p = p.Next
		} else { //有重复的数，直接删除这一段，也就是指向q
			p.Next = q
		}
	}
	return dummy.Next
}
