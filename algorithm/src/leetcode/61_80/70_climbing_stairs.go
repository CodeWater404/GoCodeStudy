package _1_80

/**
  @author: CodeWater
  @since: 2024/3/4
  @desc: 70. 爬楼梯
**/

// 从第0级往后列出方案，可以看出是一个斐波那契数列
func climbStairs(n int) int {
	a, b := 1, 1
	for n--; n > 0; n-- {
		c := a + b
		a, b = b, c
	}
	return b
}
