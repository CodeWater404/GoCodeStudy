package offer

/**
  @author: CodeWater
  @since: 2023/6/19
  @desc: 从上到下打印二叉树 III
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder3(root *TreeNode) [][]int {
	queue, res := []*TreeNode{}, [][]int{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		temp := []int{}
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if len(res)%2 == 0 {
				//当前遍历到寄数层，元素添加到队尾，这样下一次遍历队列的时候就相当于从左往右遍历二叉树
				temp = append(temp, node.Val)
			} else {
				//res中存储了奇数个元素时也就是在第偶数层，从右往左，存储到队列中时就需要存在队头
				temp = append([]int{node.Val}, temp...)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
