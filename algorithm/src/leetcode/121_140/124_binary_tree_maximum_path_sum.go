package _21_140

import "math"

/**
  @author: CodeWater
  @since: 2024/1/24
  @desc: 129. 求根到叶子节点数字之和
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans int

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans = math.MinInt32
	dfs(root)
	return ans
}

func dfs(u *TreeNode) int {
	if u == nil {
		return 0
	}
	left, right := max(0, dfs(u.Left)), max(0, dfs(u.Right))
	ans = max(ans, u.Val+left+right)
	return u.Val + max(left, right)
}
