package offer

import "math"

/**
  @author: CodeWater
  @since: 2023/6/26
  @desc: 股票的最大利润
**/
func maxProfit(prices []int) int {
	//cost记录前i天股票的最小价格； profit记录最大利润=max（前i天的最大利润，当天价格减掉前i天的最小价格）
	cost, profit := math.MaxInt32, 0
	for _, value := range prices {
		cost = min(cost, value)
		profit = max(profit, value-cost)
	}
	return profit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}