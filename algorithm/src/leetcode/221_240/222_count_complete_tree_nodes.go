package _21_240

/**
  @author: CodeWater
  @since: 2024/1/25
  @desc: 222. 完全二叉树的节点个数
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// countNodes 二分，左右子树深度不一样，说明有一边不是满二叉树
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := root.Left, root.Right
	// x、y：左右子树深度
	x, y := 1, 1
	for l != nil {
		l, x = l.Left, x+1
	}
	for r != nil {
		r, y = r.Right, y+1
	}
	// 满二叉树，直接计算：2^n - 1
	if x == y {
		return (1 << x) - 1
	}
	// 完全二叉树，左子树+当前的根节点+右子树
	return countNodes(root.Left) + 1 + countNodes(root.Right)
}
