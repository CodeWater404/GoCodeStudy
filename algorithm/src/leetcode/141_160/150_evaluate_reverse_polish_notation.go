package _41_160

import "strconv"

/**
  @author: CodeWater
  @since: 2024/1/5
  @desc: 150. 逆波兰表达式求值
**/

func evalRPN(tokens []string) int {
	var stk []int
	for _, s := range tokens {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			b := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			a := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if s == "+" {
				a += b
			} else if s == "-" {
				a -= b
			} else if s == "*" {
				a *= b
			} else {
				a /= b
			}
			stk = append(stk, a)
		} else {
			num, _ := strconv.Atoi(s)
			stk = append(stk, num)
		}
	}
	return stk[0]
}
