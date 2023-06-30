package offer

/**
  @author: CodeWater
  @since: 2023/6/30
  @desc: 数字序列中某一位的数字
**/
func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		digit += 1
		start *= 10
		count = digit * start * 9

	}
	num := start + (n-1)/digit
	numStr := strconv.FormatInt(int64(num), 10)
	digitChar := numStr[(n-1)%digit]
	res, _ := strconv.Atoi(string(digitChar))
	return res
}
