package _1_40

import "fmt"

/**
  @author: CodeWater
  @since: 2024/1/9
  @desc: 25. K 个一组翻转链表
**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	for p := dummy; ; {
		q := p
		//q指向k个节点的右边一个一个节点
		for i := 0; i < k && q != nil; i++ {
			q = q.Next
		}
		//q不空意味着范围内满足k个节点
		if q == nil {
			break
		}
		a := p.Next
		b := a.Next
		//反转k个范围的节点
		for i := 0; i < k-1; i++ {
			c := b.Next            //两个节点交换，保存下第二个节点指向的第三个节点
			b.Next, a, b = a, b, c //ab往后顺移
		}
		//改变k范围左右边界的指向
		c := p.Next
		p.Next, c.Next = a, b
		p = c
		fmt.Printf("%#+v\n", p)
	}
	return dummy.Next
}
