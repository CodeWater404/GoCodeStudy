package offer

/**
  @author: CodeWater
  @since: 2023/6/21
  @desc: 二叉树的深度
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//此树的深度和其左（右）子树的深度之间的关系。显然，此树的深度 等于 左子树的深度 与 右子树的深度 中的 最大值 +1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//递归理解：每一层都要看左右子树的深度，相当于大问题划分小问题
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
