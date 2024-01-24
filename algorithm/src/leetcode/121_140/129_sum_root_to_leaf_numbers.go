package _21_140

/**
  @author: CodeWater
  @since: 2024/1/24
  @desc: 129. 求根到叶子节点数字之和
**/

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func sumNumbers(root *TreeNode) int {
	ans = 0
	if root != nil {
		dfs(root, 0)
	}
	return ans
}

func dfs(root *TreeNode, number int) {
	number = number*10 + root.Val
	if root.Left == nil && root.Right == nil {
		ans += number
	}
	if root.Left != nil {
		dfs(root.Left, number)
	}
	if root.Right != nil {
		dfs(root.Right, number)
	}
}
