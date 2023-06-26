package offer

/**
  @author: CodeWater
  @since: 2023/6/26
  @desc: 青蛙跳台阶问题
**/

func numWays(n int) int {
	a, b := 1, 1
	var sum int
	for i := 0; i < n; i++ {
		sum = (a + b) % int(1e9+7)
		a = b
		b = sum
	}
	return a
}
