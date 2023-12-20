package __datastructure

import (
	"fmt"
	"unicode"
)

/**
  @author: CodeWater
  @since: 2023/12/21
  @desc: 表达式求值
**/

var (
	num []int  // 存储数字的栈
	op  []byte // 存储操作符的栈
)

// eval 计算一个操作符对应的表达式
func eval() {
	b := num[len(num)-1]
	num = num[:len(num)-1]
	a := num[len(num)-1]
	num = num[:len(num)-1]
	c := op[len(op)-1]
	op = op[:len(op)-1]
	var x int
	if c == '+' {
		x = a + b
	} else if c == '-' {
		x = a - b
	} else if c == '*' {
		x = a * b
	} else {
		x = a / b
	}
	num = append(num, x)
}

func main() {
	// 定义操作符的优先级
	pr := map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2}
	var str string
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		c := str[i]
		//判断当前是不是数字
		if unicode.IsDigit(rune(c)) {
			//x计算出字符串表达的数字
			x, j := 0, i
			// 当前数字不止一位
			for ; j < len(str) && unicode.IsDigit(rune(str[j])); j++ {
				x = x*10 + int(str[j]-'0') // uint8转为int
			}
			// i回到数字最后一位,下一次移动由for循环完成
			i = j - 1
			num = append(num, x)
		} else if c == '(' {
			op = append(op, c)
		} else if c == ')' {
			// 计算一个括号的表达式
			for op[len(op)-1] != '(' {
				eval()
			}
			// 弹出（
			op = op[:len(op)-1]
		} else {
			//当前字符是操作符的情况，并且op还有操作符且不是（且优先级大于当前操作符优先级
			for len(op) > 0 && op[len(op)-1] != '(' && pr[op[len(op)-1]] >= pr[c] {
				// 也就是说，在遇到操作符的时候，先计算op栈顶优先级大的，
				//前面的数字肯定已经完全入栈了，当前的-+依赖前面的计算
				eval()
			}
			// op栈顶优先级小则入栈，留着后面计算
			op = append(op, c)
		}
	}
	// 再次查看op中是否还有操作符，不用担心 （，在上一轮循坏就已经消化完了
	for len(op) > 0 {
		eval()
	}
	fmt.Println(num[len(num)-1])
}
