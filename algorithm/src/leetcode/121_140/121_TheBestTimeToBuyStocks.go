package _21_140

import "math"

/**
  @author: CodeWater
  @since: 2023/10/20
  @desc:买股票的最佳时机
**/
func maxProfit(prices []int) int {
	res := 0
	/*
	   一次遍历，res维护答案，minp维护从0开始到当前位置的股票价格最小值。
	   每次往后遍历的时候，更新res，对比当前股票价格减掉前面最小值的差和res谁更大；minp也比较当前元素谁更小
	*/
	for i, minp := 0, math.MaxInt; i < len(prices); i++ {
		res = max(res, prices[i]-minp)
		minp = min(minp, prices[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
