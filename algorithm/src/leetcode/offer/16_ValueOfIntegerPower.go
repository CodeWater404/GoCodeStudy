package offer

/**
  @author: CodeWater
  @since: 2023/6/23
  @desc: 数值的整数次方
**/
func myPow(x float64, n int) float64 {
	//x为0时直接返回0，防止下面1/x报错
	if x == 0 {
		return 0
	}
	b, res := n, 1.0
	//指数为小数时的转换
	if b < 0 {
		x = 1 / x
		b = -b
	}
	//到这里全是正数，直接进行计算
	for b > 0 {
		/*x的n次方= （x的平方）n/2次方； n/2有可能奇数有可能偶数；奇数的话需要多乘一个x，除以2是因为n每次进行左移一位，所以这么计算*/
		//指数b最右边位置上是1，
		if (b & 1) == 1 {
			res *= x
		}
		//指数最右边不是1，说明答案不需要乘
		x *= x
		b >>= 1
	}
	return res
}
