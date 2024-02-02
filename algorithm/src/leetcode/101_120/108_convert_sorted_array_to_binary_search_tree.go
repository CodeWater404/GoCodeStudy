package _01_120

/**
  @author: CodeWater
  @since: 2024/2/2
  @desc: 将有序数组转换为二叉搜索树
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	root := &TreeNode{Val: nums[mid]}
	root.Left, root.Right = build(nums, l, mid-1), build(nums, mid+1, r)
	return root
}
