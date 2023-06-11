package offer

/**
  @author: CodeWater
  @since: 2023/6/11
  @desc: 18. 删除链表的节点
**/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	prev, cur := head, head.Next
	for cur != nil && cur.Val != val {
		prev, cur = cur, cur.Next
	}
	if cur != nil {
		prev.Next = cur.Next
	}
	return head
}
