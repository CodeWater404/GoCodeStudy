package _21_540

import "math"

/**
  @author: CodeWater
  @since: 2024/1/28
  @desc: 二叉搜索树的最小绝对差
**/

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	ans, last int
	is_first  bool
)

// getMinimumDifference 二叉搜索树的相邻节点差值就是最小的
func getMinimumDifference(root *TreeNode) int {
	ans, is_first = math.MaxInt32, true
	dfs(root)
	return ans
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if is_first {
		// 特判是不是第一个数，第一个数不用计算
		is_first = false
	} else {
		ans = min(ans, root.Val-last)
	}
	// 当前数存到last里面，方便下一轮计算
	last = root.Val
	dfs(root.Right)

}
