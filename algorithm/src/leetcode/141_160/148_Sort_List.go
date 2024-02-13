package _41_160

/**
  @author: CodeWater
  @since: 2024/2/13
  @desc: 148. 排序链表
**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// sortList 归并排序的从底向上迭代
func sortList(head *ListNode) *ListNode {
	n := 0
	for p := head; p != nil; p = p.Next { //计算链表长度
		n++
	}
	for i := 1; i < n; i *= 2 { //i要小于n，等于n时已经排好序
		dummy := &ListNode{Val: -1}
		cur := dummy
		// 每次枚举两个区间
		for j := 1; j <= n; j += i * 2 { // j更新时是往后错开两个区间
			p := head // 第一个指针p和第二个指针q；这里把head存为每次比较的头节点
			q := p
			for k := 0; k < i && q != nil; k++ { //q往后跳i步
				q = q.Next
			}
			o := q
			for k := 0; k < i && o != nil; k++ { // o存储下一轮两个区间比较的开头
				o = o.Next
			}
			//归并
			l, r := 0, 0
			for l < i && r < i && p != nil && q != nil {
				// 比较哪一个指针指向的值小，然后更新cur的位置
				if p.Val <= q.Val {
					cur.Next = p
					cur = cur.Next
					p = p.Next
					l++
				} else {
					cur.Next = q
					cur = cur.Next
					q = q.Next
					r++
				}
			}
			// 把两个区间内还剩余的元素补到cur后面，更新cur位置
			for l < i && p != nil {
				cur.Next = p
				cur = cur.Next
				p = p.Next
				l++
			}
			for r < i && q != nil {
				cur.Next = q
				cur = cur.Next
				q = q.Next
				r++
			}
			// head更新到下一次两个区间比较的开头
			head = o
		}
		// 每一层做完之后，把尾节点置空；头节点更新成这一层的头节点
		cur.Next = nil
		head = dummy.Next
	}
	return head
}
