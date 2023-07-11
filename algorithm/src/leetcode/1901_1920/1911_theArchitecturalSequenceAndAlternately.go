package _901_1920

/**
  @author: CodeWater
  @since: 2023/7/11
  @desc: 最大子序列交替和
**/
func maxAlternatingSum1(nums []int) int64 {
	n := len(nums)
	//f[i] 表示从前i 个元素中选出的子序列，且最后一个元素为奇数下标时的最大交替和
	//g[i] 表示从前 i 个元素中选出的子序列，且最后一个元素为偶数下标时的最大交替和
	f, g := make([]int, n+1), make([]int, n+1)
	for i, x := range nums {
		//这里i++更多的做一个初始化？如果不用，f[i-1]报错。第i个元素 nums[i−1]：
		i++
		//如果当前是奇数下标并且选取的话，那么对于f[i]来说要么不选是f[i-1];要么选那交替和就是g[i-1]-x，由偶数结尾的减去当前值（因为当前是奇数下标，所以上一个交替和是偶数结尾的g[i-1]）。 g[i]同理
		f[i] = max(f[i-1], g[i-1]-x)
		g[i] = max(g[i-1], f[i-1]+x)
	}
	return int64(max(f[n], g[n]))
}

func maxAlternatingSum2(nums []int) int64 {
	var f, g int
	//优化：因为f[i] 只和g[i]上一个f[i-1]和g[i - 1]有关，所以直接用变量替代
	for _, x := range nums {
		f, g = max(g-x, f), max(f+x, g)
	}
	return int64(max(f, g))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
