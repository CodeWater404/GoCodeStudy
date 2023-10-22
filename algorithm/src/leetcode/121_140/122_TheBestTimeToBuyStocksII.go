package _21_140

/**
  @author: CodeWater
  @since: 2023/10/22
  @desc: 买股票的最佳时机II
**/

func maxProfit(prices []int) int {
	res := 0
	for i := 0; i+1 < len(prices); i++ {
		res += max(0, prices[i+1]-prices[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
