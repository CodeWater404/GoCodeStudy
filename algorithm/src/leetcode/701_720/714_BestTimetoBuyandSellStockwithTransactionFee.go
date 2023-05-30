package _01_720

/**
  @author: CodeWater
  @since: 2023/5/30
  @desc: 714. 买卖股票的最佳时机含手续费
**/

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	/*状态dp，图画下来就是一个有权图，也就可以把题目转换成走n条边获取到的利润最大是多少
	  0状态：手上无货
	  1状态：手上有货（股票）
	  f[i][0] ：走了i条边，处于0状态的最大利润
	  f[i][1] ：走了i条边，处于1状态的最大利润
	*/
	f, res, INF := make([][]int, n+1), 0, int(1e8)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, 2)
		f[i][0] = -INF
		f[i][1] = -INF
	}
	//一开始手上无货，利润0
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		// 走到0状态有两种：上一次的0，上一次的1（1状态同理.从0到1也就是买入，所以算利润的时候还需要减去买入股票的价格以及手续费fee；相反1到0就是卖出。本题中fee算在买入和卖出中都可以，无影响）
		f[i][0] = max(f[i-1][0], f[i-1][1]+prices[i-1]) //prices[i - 1]下标从1开始，所以减1
		f[i][1] = max(f[i-1][1], f[i-1][0]-prices[i-1]-fee)
		// 最后是无货状态，所以是和0状态比较，更新
		res = max(res, f[i][0])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
