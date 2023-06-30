package offer

/**
  @author: CodeWater
  @since: 2023/6/30
  @desc: 1～n 整数中 1 出现的次数
**/
func countDigitOne(n int) int {
	digit, res := 1, 0
	high, cur, low := n/10, n%10, 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}
