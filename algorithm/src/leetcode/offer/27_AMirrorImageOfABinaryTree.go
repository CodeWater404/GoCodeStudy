package offer

/**
  @author: CodeWater
  @since: 2023/6/19
  @desc: 二叉树的镜像
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(temp)
	return root
}
