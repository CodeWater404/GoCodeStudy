package _1_60

import "math"

/**
  @author: CodeWater
  @since: 2024/3/2
  @desc: 50. Pow(x, n)
**/

// 快速幂
func myPow(x float64, n int) float64 {
	is_minus, res := n < 0, 1.0
	for k := int64(math.Abs(float64(n))); k > 0; k >>= 1 { //枚举k的最低位，然后右移删除该位
		if k&1 == 1 {
			res *= x
		}
		// 这里就是计算下一步res要乘上的x
		x *= x
	}
	if is_minus { // n是负数的情况，取倒数
		res = 1 / res
	}
	return res
}
