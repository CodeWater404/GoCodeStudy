package offer

/**
  @author: CodeWater
  @since: 2023/6/11
  @desc:22. 链表中倒数第k个节点
**/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	//两个指针保持距离k即可
	pre, cur := head, head
	for k != 0 {
		cur = cur.Next
		k--
	}
	for cur != nil {
		pre = pre.Next
		cur = cur.Next
	}
	return pre
}
