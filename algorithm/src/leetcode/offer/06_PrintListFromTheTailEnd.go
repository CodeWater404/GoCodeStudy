package offer

/**
  @author: CodeWater
  @since: 2023/6/9
  @desc: 从头到尾打印链表
**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// go没有现成的栈类型
func reversePrint(head *ListNode) []int {
	arr := []int{}
	current := head
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	for i, j := 0, len(arr)-1; i < j; i++ {
		arr[i], arr[j] = arr[j], arr[i]
		j--
	}
	return arr
}
