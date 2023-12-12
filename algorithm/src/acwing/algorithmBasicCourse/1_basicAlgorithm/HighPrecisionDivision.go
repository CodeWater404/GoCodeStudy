package __basicAlgorithm

import (
	"fmt"
	"math/big"
)

/**
  @author: CodeWater
  @since: 2023/8/24
  @desc: 高精度除法
**/

func div(a []int, b int, r *int) []int {
	var c []int
	*r = 0
	for i := len(a) - 1; i >= 0; i-- {
		*r = *r*10 + a[i]
		c = append(c, *r/b)
		*r %= b
	}
	//var temp []int
	//for i := len(c) - 1; i >= 0; i-- {
	//	temp = append(temp, c[i])
	//}
	////反转c
	//c = temp

	// 因为是商保存的时候0下标是从高位到低位开始，所以要反转一下（为了和其他三种高精度保持一致的存储）
	//去除前导0======》优化写法
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}

	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c
}

// ==================== 库函数 ====================
func div2() {
	var a, b big.Int
	fmt.Scan(&a, &b)
	// divmod其实会返回两个数商和余数，商赋值给第一个参数，余数赋值给第三个参数，这里是简写
	// 正常是这样：s , r = a.DivMod(&a , &b , &r)
	a.DivMod(&a, &b, &b)
	fmt.Printf("%d\n%d", &a, &b)
}

func main() {
	var (
		a string
		b int
		A []int
		r int
	)
	fmt.Scan(&a, &b)
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}

	c := div(A, b, &r)

	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
	fmt.Printf("\n%d", r)

}
