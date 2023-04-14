package fib

/**
  @author: CodeWater
  @since: 2023/4/14
  @desc: $
**/

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
