package _21_340

/**
  @author: CodeWater
  @since: 2024/3/6
  @desc: 322. 零钱兑换
**/

func coinChange(coins []int, m int) int {
	f := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		f[i] = 1e8
	}
	f[0] = 0
	for _, v := range coins {
		for j := v; j <= m; j++ {
			f[j] = min(f[j], f[j-v]+1)
		}
	}
	if f[m] == 1e8 {
		return -1
	}
	return f[m]
}
