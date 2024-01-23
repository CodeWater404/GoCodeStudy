package _01_120

/**
  @author: CodeWater
  @since: 2024/1/23
  @desc: 112. 路径总和
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return root.Left != nil && hasPathSum(root.Left, sum) || root.Right != nil && hasPathSum(root.Right, sum)
}
