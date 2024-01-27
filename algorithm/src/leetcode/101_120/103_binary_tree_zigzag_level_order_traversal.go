package _01_120

/**
  @author: CodeWater
  @since: 2024/1/27
  @desc: 103. 二叉树的锯齿形层序遍历
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q, cnt := make([]*TreeNode, 0), 0
	q = append(q, root)
	for len(q) > 0 {
		level, n := make([]int, 0), len(q)
		for ; n > 0; n-- {
			t := q[0]
			level, q = append(level, t.Val), q[1:]
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		// 偶数层时就反转一次
		cnt++
		if cnt%2 == 0 {
			reverse(level)
		}
		res = append(res, level)
	}
	return res
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
