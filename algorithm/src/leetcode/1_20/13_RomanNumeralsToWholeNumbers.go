package __20

/**
  @author: CodeWater
  @since: 2023/10/28
  @desc: 罗马数字转整数
**/
func romanToInt(s string) int {
	// hash映射字符和数值
	hash := make(map[byte]int)
	hash['I'] = 1
	hash['V'] = 5
	hash['X'] = 10
	hash['L'] = 50
	hash['C'] = 100
	hash['D'] = 500
	hash['M'] = 1000
	// 更方便的声明： hash := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	res := 0
	// 从前往后遍历字符，特判：当前字符代表的数值比后一个要小的时候就减掉当前字符代表的数值
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && hash[s[i]] < hash[s[i+1]] {
			res -= hash[s[i]]
		} else {
			res += hash[s[i]]
		}
	}
	return res
}
