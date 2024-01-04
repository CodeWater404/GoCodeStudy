package __20

import "math"

/**
  @author: CodeWater
  @since: 2024/1/4
  @desc: 20. 有效的括号
**/

// 这种类型的"(([))]"是false，遍历的时候当前括号只考虑和栈顶的比较
func isValid(s string) bool {
	stk, tt := make([]byte, 10010), -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			tt++
			stk[tt] = s[i]
		} else {
			//匹配的一对括号的ascii码值绝对值小于2，记不住可以直接if，这里省代码
			if tt >= 0 && math.Abs(float64(stk[tt])-float64(s[i])) <= 2 {
				tt--
			} else {
				return false
			}
		}
	}
	return tt < 0

}
