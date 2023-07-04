package _661_2680

import "sort"

/**
  @author: CodeWater
  @since: 2023/7/4
  @desc: 矩阵中的和
**/

func matrixSum(nums [][]int) int {
	res := 0
	m := len(nums)
	n := len(nums[0])
	//排序，保证每次内层选出最大的
	for i := 0; i < m; i++ {
		sort.Ints(nums[i])
	}
	//注意是先遍历内层，这样每次取到内层最大一个。
	for j := 0; j < n; j++ {
		maxVal := 0
		for i := 0; i < m; i++ {
			if nums[i][j] > maxVal {
				maxVal = nums[i][j]
			}
		}
		res += maxVal
	}
	return res
}
