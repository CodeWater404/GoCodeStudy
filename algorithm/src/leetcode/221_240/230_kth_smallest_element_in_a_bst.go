package _21_240

/**
  @author: CodeWater
  @since: 2024/1/28
  @desc: 二叉搜索树中第K小的元素
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ans, kk int

func kthSmallest(root *TreeNode, k int) int {
	kk = k
	dfs(root)
	return ans
}

// dfs bool是为了提前返回，省掉后面的遍历
func dfs(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if dfs(root.Left) {
		return true
	}
	kk--
	if kk == 0 {
		ans = root.Val
		return true
	}
	return dfs(root.Right)
}
