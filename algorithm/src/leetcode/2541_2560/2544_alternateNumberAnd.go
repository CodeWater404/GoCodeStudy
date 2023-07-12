package _541_2560

/**
  @author: CodeWater
  @since: 2023/7/12
  @desc: 交替数字和
**/
func alternateDigitSum1(n int) int {
	length, temp := 0, n
	//求出n多少位，奇数位个位开始为正号，偶数负
	for temp != 0 {
		length++
		temp /= 10
	}
	res := 0
	temp = length
	for i := 0; i < length; i++ {
		num := n % 10
		if temp%2 == 0 {
			num = -num
		}
		res += num
		//更新
		n /= 10
		temp--
	}
	return res
}

func alternateDigitSum2(n int) int {
	res, sign := 0, 1
	for n > 0 {
		res += n % 10 * sign
		sign = -sign
		n /= 10
	}
	return -sign * res
}
