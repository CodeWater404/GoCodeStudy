package offer

/**
  @author: CodeWater
  @since: 2023/6/11
  @desc: 25. 合并两个排序的链表
**/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = &ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = l2
			cur, l2 = l2, l2.Next
		} else {
			cur.Next = l1
			cur, l1 = l1, l1.Next
		}
		//上面的cur节点的更新其实可以在这里优化一下写法
		//cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}
