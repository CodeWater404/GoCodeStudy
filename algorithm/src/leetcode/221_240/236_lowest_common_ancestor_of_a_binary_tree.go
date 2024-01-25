package _21_240

/**
  @author: CodeWater
  @since: 2024/1/26
  @desc: 236. 二叉树的最近公共祖先
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans = nil
	dfs(root, p, q)
	return ans
}

// dfs 返回值表示的
func dfs(root, p, q *TreeNode) int {
	if root == nil {
		return 0
	}
	//二进制个位表示有p，十位表示有q，都有的时候就是11，也就是3
	state := dfs(root.Left, p, q)
	if root == p {
		state |= 1
	} else if root == q {
		state |= 2
	}
	state |= dfs(root.Right, p, q)
	// 当前节点下有pq，而且是第一个公共祖先节点（ans为空）
	if state == 3 && ans == nil {
		ans = root
	}
	return state
}
