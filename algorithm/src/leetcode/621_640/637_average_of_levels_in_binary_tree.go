package _21_640

/**
  @author: CodeWater
  @since: 2024/1/26
  @desc: 二叉树的层平均值
**/

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	q, res := make([]*TreeNode, 0), make([]float64, 0)
	if root == nil {
		return res
	}
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		sum := 0
		for i := 0; i < n; i++ {
			t := q[0]
			q, sum = q[1:], sum+t.Val
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
			if i == n-1 {
				res = append(res, float64(sum)/float64(n))
			}
		}
	}
	return res
}
