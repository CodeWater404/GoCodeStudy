package _01_120

/**
  @author: CodeWater
  @since: 2024/1/20
  @desc: 101. 对称二叉树
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

// dfs p左子树 ，q右子树
func dfs(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
}
