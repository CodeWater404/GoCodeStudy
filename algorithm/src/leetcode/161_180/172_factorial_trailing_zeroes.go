package _61_180

/**
  @author: CodeWater
  @since: 2024/2/28
  @desc: 172. 阶乘后的零
**/

/*
要看n！末尾有多少个0，就要看他是10的多少倍。10=2x5，也就是看2和5出现了多少对，一对可以乘起来变成10，
也就是最后看2和5最少有多少个因子再n！的里面。
*/
func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n /= 5
	}
	return res
}
