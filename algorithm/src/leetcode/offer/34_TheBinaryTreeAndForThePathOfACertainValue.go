package offer

/**
  @author: CodeWater
  @since: 2023/6/20
  @desc: 二叉树中和为某一值的路径
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
	res  [][]int
	path []int
)

func pathSum(root *TreeNode, target int) [][]int {
	//一定要初始化，不初始化这次是空树的话，那么就会保留上一次的答案然后return了
	res = [][]int{}
	path = []int{}
	recur2(root, target)
	return res
}

func recur2(root *TreeNode, tar int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	tar = tar - root.Val
	if tar == 0 && root.Left == nil && root.Right == nil {
		//append实际上append的是path的地址，如果直接append，之后path的值发生改变的话res里面的值也会变,所以这里弄个空数组然后path添加进去。[]int(nil) 表示一个空的整数切片，相当于 []int{}。
		//所以不能这么写：res = append(res , path...)
		res = append(res, append([]int(nil), path...))
	}
	recur2(root.Left, tar)
	recur2(root.Right, tar)
	path = path[:len(path)-1]

}
