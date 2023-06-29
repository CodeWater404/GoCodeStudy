package offer

/**
  @author: CodeWater
  @since: 2023/6/29
  @desc: 剪绳子 II
**/
func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	b, p := n%3, int(1e9+7)
	rem, x := 1, 3
	for a := n/3 - 1; a > 0; a /= 2 {
		if a%2 == 1 {
			rem = (rem * x) % p
		}
		x = (x * x) % p
	}
	if b == 0 {
		return int(rem * 3 % p)
	}
	if b == 1 {
		return int(rem * 4 % p)
	}
	return int(rem * 6 % p)
}
