package __basicAlgorithm

import "fmt"

/**
  @author: CodeWater
  @since: 2023/12/16
  @desc: 二进制中1的个数(lowbit)
**/

// lowbit 返回从地位开始的第一位1，原理：-x在计算机中用补码表示，也就是x反码+1，所以
// x&-x之后就能得到最低位的一
func lowbit(x int) int {
	return x & -x
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var m int
		res := 0
		fmt.Scanf("%d", &m)
		for m > 0 {
			// 每次减掉最低位的1，每减一次就说明1出现一次，res++
			m -= lowbit(m)
			res++
		}
		fmt.Printf("%d ", res)
	}
}
