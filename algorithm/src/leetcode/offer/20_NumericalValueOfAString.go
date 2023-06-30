package offer

/**
  @author: CodeWater
  @since: 2023/6/30
  @desc: 表示数值的字符串
**/
func isNumber(s string) bool {
	states := []map[rune]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4}, // 0.
		{'d': 2, '.': 4},                 // 1.
		{'d': 2, '.': 3, 'e': 5, ' ': 8}, // 2.
		{'d': 3, 'e': 5, ' ': 8},         // 3.
		{'d': 3},                         // 4.
		{'s': 6, 'd': 7},                 // 5.
		{'d': 7},                         // 6.
		{'d': 7, ' ': 8},                 // 7.
		{' ': 8},                         // 8.
	}

	p := 0
	for _, c := range s {
		var t rune
		if c >= '0' && c <= '9' {
			t = 'd'
		} else if c == '+' || c == '-' {
			t = 's'
		} else if c == 'e' || c == 'E' {
			t = 'e'
		} else if c == '.' || c == ' ' {
			t = c
		} else {
			t = '?'
		}

		if _, ok := states[p][t]; !ok {
			return false
		}
		p = states[p][t]
	}

	return p == 2 || p == 3 || p == 7 || p == 8
}
