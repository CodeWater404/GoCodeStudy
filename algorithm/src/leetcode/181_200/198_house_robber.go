package _81_200

/**
  @author: CodeWater
  @since: 2024/3/4
  @desc: 198. 打家劫舍
**/

func rob(nums []int) int {
	n := len(nums)
	//f[i]:选择第i个点；g[i]:不选第i个点
	f, g := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		//第i个点选择的话，那么第i-1个点就不能选，所以是前面最大值g[i-1]+当前第i个点nums[i-1]
		f[i] = g[i-1] + nums[i-1]
		//第i个点不选的话，那么第i-1个点就可以选择f[i-1]或者不选g[i-1]
		g[i] = max(f[i-1], g[i-1])
	}
	return max(f[n], g[n])
}
