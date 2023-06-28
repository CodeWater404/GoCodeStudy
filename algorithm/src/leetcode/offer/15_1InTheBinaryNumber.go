package offer

/**
  @author: CodeWater
  @since: 2023/6/28
  @desc: 二进制中1的个数
**/
func hammingWeight(num uint32) int {
	sum := 0
	for num != 0 {
		// if num & 1 == 1 {
		//     sum += 1
		// }
		//num = num >> 1

		//简洁写法
		sum += int(num & 1)
		num >>= 1
	}
	return sum
}