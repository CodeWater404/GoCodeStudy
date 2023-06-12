package offer

/**
  @author: CodeWater
  @since: 2023/6/12
  @desc: 两个链表的第一个公共节点
**/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	// 这里交叉走其实已经考虑了不相交的情况，就是ab同时为空
	for a != b {
		//go没有三目运算符，用if
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}
