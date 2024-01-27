package _01_120

/**
  @author: CodeWater
  @since: 2024/1/27
  @desc: 二叉树的层序遍历
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		level := make([]int, 0)
		for i := 0; i < n; i++ {
			node := q[0]
			level, q = append(level, node.Val), q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
