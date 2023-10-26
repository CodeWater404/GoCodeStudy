package _21_240

/**
  @author: CodeWater
  @since: 2023/10/26
  @desc: 除自身以外数组的乘积
**/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	p := make([]int, n)
	for i := range p {
		p[i] = 1
	}
	for i := 1; i < n; i++ {
		p[i] = p[i-1] * nums[i-1]
	}
	for i, s := n-1, 1; i >= 0; i-- {
		p[i] *= s
		s *= nums[i]
	}
	return p
}
