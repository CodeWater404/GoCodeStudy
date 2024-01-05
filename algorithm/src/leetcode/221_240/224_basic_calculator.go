package _21_240

import "unicode"

/**
  @author: CodeWater
  @since: 2024/1/5
  @desc: 224. 基本计算器
**/

var (
	num []int
	op  []byte
)

func calculate(rs string) int {
	s := ""
	for _, c := range rs {
		if c != ' ' {
			s += string(c)
		}
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if unicode.IsDigit(rune(c)) {
			x, j := 0, i
			for j < len(s) && unicode.IsDigit(rune(s[j])) {
				x = x*10 + int(s[j]-'0')
				j++
			}
			i = j - 1
			num = append(num, x)
		} else if c == '(' {
			op = append(op, c)
		} else if c == ')' {
			for op[len(op)-1] != '(' {
				eval()
			}
			op = op[:len(op)-1]
		} else {
			//本题只有加减，不用考虑运算符的优先级情况
			//处理这种情况"1-(     -2)"，num加一个数进去方便eval计算
			if i == 0 || s[i-1] == '(' || s[i-1] == '+' || s[i-1] == '-' {
				num = append(num, 0)
			}
			for len(op) > 0 && op[len(op)-1] != '(' {
				eval()
			}
			op = append(op, c)
		}
	}
	for len(op) > 0 {
		eval()
	}
	return num[len(num)-1]
}

// 注意如果用参数需要用指针，虽然切片是引用类型
func eval() {
	b := num[len(num)-1]
	num = num[:len(num)-1]
	a := num[len(num)-1]
	num = num[:len(num)-1]
	c := op[len(op)-1]
	op = op[:len(op)-1]
	r := 0
	if c == '+' {
		r = a + b
	} else {
		r = a - b
	}
	num = append(num, r)
}
