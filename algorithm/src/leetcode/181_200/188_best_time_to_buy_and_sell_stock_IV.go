package _81_200

/**
  @author: CodeWater
  @since: 2024/3/30
  @desc: 188. 买卖股票的最佳时机 IV
**/

var (
	f [10001]int
	g [10001]int
)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if k > n/2 {
		res := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				res += prices[i] - prices[i-1]
			}
		}
		return res
	}
	for i := 0; i < 10001; i++ {
		f[i], g[i] = -0x3f, -0x3f
	}
	f[0] = 0
	res := 0
	for i := 1; i <= n; i++ {
		for j := k; j >= 0; j-- {
			g[j] = max(g[j], f[j]-prices[i-1])
			if j > 0 {
				f[j] = max(f[j], g[j-1]+prices[i-1])
			}
		}
	}
	for i := 1; i <= k; i++ {
		res = max(res, f[i])
	}
	return res
}
