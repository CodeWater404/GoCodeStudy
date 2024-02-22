package _81_200

/**
  @author: CodeWater
  @since: 2024/2/22
  @desc: 191. 位1的个数
**/

// 方法一： 依次枚举每一位
func hammingWeight1(n uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		if n>>i&1 == 1 {
			res++
		}
		//可以优化成：res += n >> i & 1
	}
	return res
}

// low bit操作：返回x的最后一位1.
// 具体操作就是每次n减去low bit就得到下一轮的数，同时记录这一次的1到答案中
func hammingWeight2(n uint32) int {
	res := 0
	for n > 0 {
		// n和-n相与就得到最低位的一个1
		// ps:n&（-n + 1）等价于n & (-n)
		n -= n & (-n)
		res++ //记录减的的个数，就是1的个数
	}
	return res
}
