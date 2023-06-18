package offer

/**
  @author: CodeWater
  @since: 2023/6/18
  @desc: 第一个只出现一次的字符
**/

func firstUniqChar(s string) byte {
	hash := make(map[byte]int)
	sc := []byte(s)

	for _, c := range sc {
		hash[c]++
	}

	//这里再遍历一边原字符串数组就是为了保持顺序性第一个只出现一次
	for _, c := range sc {
		if hash[c] == 1 {
			return c
		}
	}
	return ' '
}
