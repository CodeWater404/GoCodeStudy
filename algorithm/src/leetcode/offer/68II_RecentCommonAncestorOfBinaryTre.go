package offer

/**
  @author: CodeWater
  @since: 2023/6/22
  @desc: 二叉树的最近公共祖先
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//注意本题是二叉树！！！！
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
