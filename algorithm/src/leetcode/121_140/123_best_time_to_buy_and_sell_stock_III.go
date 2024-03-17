package _21_140

import "math"

/**
  @author: CodeWater
  @since: 2024/3/17
  @desc: 123. 买卖股票的最佳时机 III
**/

func maxProfit(prices []int) int {
	n := len(prices)
	f := make([]int, n+2)
	//预处理前缀，i从1开始，所以第i天就是f(i-1)
	for i, minp := 1, math.MaxInt32; i <= n; i++ {
		//f(i-1):第i天不卖出，比较第i天卖出：f(i-1)和i-1减掉minp的最大值
		f[i] = max(f[i-1], prices[i-1]-minp)
		//minp：记录第i天之前的最小价格
		minp = min(minp, prices[i-1])
	}

	res := 0
	//处理后缀
	for i, maxp := n, 0; i > 0; i-- {
		//res和第i天卖出的情况（maxp-prices[i-1]）加上1到i-1之间的最值f(i-1)。i时分界点
		res = max(res, maxp-prices[i-1]+f[i-1])
		maxp = max(maxp, prices[i-1])
	}
	return res
}
