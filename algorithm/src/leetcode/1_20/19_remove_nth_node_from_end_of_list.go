package __20

/**
  @author: CodeWater
  @since: 2024/1/10
  @desc: 19. 删除链表的倒数第 N 个结点
**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy, m := &ListNode{Val: -1}, 0
	dummy.Next = head
	for p := dummy; p != nil; p = p.Next {
		m++
	}
	p := dummy
	for i := 0; i < m-n-1; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}
