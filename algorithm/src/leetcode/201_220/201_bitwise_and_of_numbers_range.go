package _01_220

/*
*

	@author: CodeWater
	@since: 2024/2/26
	@desc: 201. 数字范围按位

*
*/
func rangeBitwiseAnd(m int, n int) int {
	res := 0
	for i := 30; i >= 0; i-- {
		// m和n的第i位不一样
		if (m >> i & 1) != (n >> i & 1) {
			break
		}
		//m第i位为1，加入到答案中
		if (m >> i & 1) > 0 {
			res += 1 << i
		}
	}
	return res
}
