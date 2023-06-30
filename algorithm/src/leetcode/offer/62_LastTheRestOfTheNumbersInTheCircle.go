package offer

/**
  @author: CodeWater
  @since: 2023/6/30
  @desc:  圆圈中最后剩下的数字
**/
func lastRemaining(n int, m int) int {
	//dp[1]:一个状态就是f(1.m)值是0（n为1，不管m多少永远得到的只能是0）
	x := 0
	//dp[i] = ( dp[i - 1] + m ) % i。 dp方程从1开始，1的值已知，所以从2开始遍历
	for i := 2; i <= n; i++ {
		//推算下一个状态
		x = (x + m) % i
	}
	return x
}
