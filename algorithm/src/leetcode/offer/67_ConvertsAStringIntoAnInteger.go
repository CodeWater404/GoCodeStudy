package offer

/**
  @author: CodeWater
  @since: 2023/6/8
  @desc: $
**/
func strToInt(str string) int {
	res, border := 0, math.MaxInt32/10
	i, sign, length := 0, 1, len(str)
	if length == 0 {
		return 0
	}
	for str[i] == ' ' {
		i++
		if i == length {
			return 0
		}
	}

	if str[i] == '-' {
		sign = -1
	}
	if str[i] == '-' || str[i] == '+' {
		i++
	}
	for j := i; j < length; j++ {
		// 非数字
		if str[j] < '0' || str[j] > '9' {
			break
		}
		digit := str[j] - '0'
		if res > border || res == border && digit > 7 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		res = res*10 + int(digit)
	}
	return sign * res
}
