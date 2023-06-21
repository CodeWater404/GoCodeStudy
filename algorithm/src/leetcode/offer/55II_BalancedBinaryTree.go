package offer

/**
  @author: CodeWater
  @since: 2023/6/21
  @desc: 平衡二叉树
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return recur(root) != -1
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	//-1左子树的左右子树高度在1之间
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	if right == -1 {
		return -1
	}
	var diff int
	//比较两个左右子树的高度差
	if abs(left, right) < 2 {
		diff = max(left, right) + 1
	} else {
		diff = -1
	}
	return diff
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
