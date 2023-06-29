package offer

/**
  @author: CodeWater
  @since: 2023/6/29
  @desc: 剪绳子 I
**/
func cuttingRope1(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return pow(3, a)
	}
	if b == 1 {
		return pow(3, a-1) * 4
	}
	return pow(3, a) * 2
}

func pow(x, y int) int {
	res := 1
	for y != 0 {
		res *= x
		y--
	}
	return res
}
