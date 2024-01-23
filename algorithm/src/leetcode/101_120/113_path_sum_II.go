package _01_120

/**
  @author: CodeWater
  @since: 2024/1/23
  @desc: 113. 路径总和 II
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	ans  [][]int
	path []int
)

func pathSum(root *TreeNode, sum int) [][]int {
	//不要在全局初始化，在全局初始化里面保存的都是第一次的结果
	ans = make([][]int, 0)
	path = make([]int, 0)
	if root != nil {
		dfs(root, sum)
	}
	return ans
}

func dfs(root *TreeNode, sum int) {
	path = append(path, root.Val)
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			// 注意：这里需要创建一个新的切片，以避免 path 切片在后续递归中被修改影响结果
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
		}
	} else {
		if root.Left != nil {
			dfs(root.Left, sum)
		}
		if root.Right != nil {
			dfs(root.Right, sum)
		}
	}
	path = path[:len(path)-1]
}
