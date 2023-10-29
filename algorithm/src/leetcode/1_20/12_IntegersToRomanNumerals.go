package __20

/**
  @author: CodeWater
  @since: 2023/10/29
  @desc: 整数转罗马数字
**/

func intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
	}
	reps := []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
	}
	res := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			res += reps[i]
		}
	}
	return res
}
