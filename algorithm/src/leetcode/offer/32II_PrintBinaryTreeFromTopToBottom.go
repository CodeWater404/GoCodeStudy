package offer

/**
  @author: CodeWater
  @since: 2023/6/18
  @desc: 从上到下打印二叉树 II
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	var (
		queue []*TreeNode
		res   [][]int
	)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		var temp []int
		//这道题标志一层关键就在len（queue），在进入内层for循环的时候遍历次数固定了；但是len在更新，
		//退出内层for的时候下一层的遍历次数也就再次确定。从而确保每一层的遍历次数
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
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
