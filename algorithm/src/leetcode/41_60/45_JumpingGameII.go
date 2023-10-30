package _1_60

/**
  @author: CodeWater
  @since: 2023/10/24
  @desc: 跳跃游戏II
**/
func jump(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j+nums[j] < i {
			j++
		}
		f[i] = f[j] + 1
	}
	return f[n-1]
}
