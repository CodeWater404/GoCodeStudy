package offer

/**
  @author: CodeWater
  @since: 2023/6/26
  @desc: 斐波那契数列
**/

func fib(n int) int {
	a, b, sum := 0, 1, 0
	for i := 0; i < n; i++ {
		sum = (a + b) % int(1e9+7)
		a = b
		b = sum
	}

	return a
}
