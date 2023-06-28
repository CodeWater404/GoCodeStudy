package offer

/**
  @author: CodeWater
  @since: 2023/6/28
  @desc: 数组中数字出现的次数
**/
func singleNumbers(nums []int) []int {
	//xy只出现一次的数，n是xy数组异或之后只剩下xy的结果，m是xy二进制位第一次不同的结果
	x, y, n, m := 0, 0, 0, 1
	//相同的数异或还是0，这里异或完，n就是不同的两个数xy疑惑的结果
	for _, num := range nums {
		n ^= num
	}
	//这里通过与运算找到xy二进制位第一个不同位
	for (n & m) == 0 {
		m <<= 1
	}
	//
	for _, num := range nums {
		//通过与运算，把nums数组分成两组，分组的规则是和 m 相与是否为1的，为1的在一组不断异或；不为1的在另外一组不断异或，遍历完数组之后就可以得到两个只出现一次的数xy
		if (num & m) != 0 {
			x ^= num
		} else {
			y ^= num
		}
	}
	return []int{x, y}
}
