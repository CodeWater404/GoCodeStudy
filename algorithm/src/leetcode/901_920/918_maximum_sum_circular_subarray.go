package _01_920

import "math"

/**
  @author: CodeWater
  @since: 2024/2/14
  @desc: 918. 环形子数组的最大和
**/

// 线段上连续区间求和用前缀和，环上的问题可以破环成链
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	// 前缀和
	s := make([]int, n*2+1)
	for i := 1; i <= n*2; i++ { //从环任意一处断开，都可以在0-2n线段上找到对应的一处区间
		s[i] = s[i-1] + nums[(i-1)%n] //然后这个0-2n线段用前缀和预处理
	}
	// Si - Sj要最大，那么Sj要最小，Sj最小就是滑动窗口
	q := make([]int, 0)
	q = append(q, 0) // s[0]
	res := math.MinInt32
	for i := 1; i <= n*2; i++ {
		for len(q) > 0 && i-q[0] > n { // 队头不空并且队头元素出界滑动窗口，删除队头
			q = q[1:]
		}
		res = max(res, s[i]-s[q[0]])
		for len(q) > 0 && s[q[len(q)-1]] >= s[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i) //i加入到单调队列中
	}
	return res
}
