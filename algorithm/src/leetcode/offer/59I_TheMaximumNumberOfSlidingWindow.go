package offer

/**
  @author: CodeWater
  @since: 2023/6/14
  @desc: 滑动窗口的最大值
**/
func maxSlidingWindow(nums []int, k int) []int {
	//q优先队列，这里存储顺序是递减的
	q := []int{}
	//push函数，始终保持q存储的正确性，另外存储的是下标
	push := func(i int) {
		//q队列不空，数组当前元素比队列队尾元素大就删除队尾，
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		//确保数组当前元素已经比队尾的头小了，加入到队尾
		q = append(q, i)
	}

	// 先初始化一下优先队列q
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	//答案长度为1，最大容量为n-k+1
	ans := make([]int, 1, n-k+1)
	//因为前面已经对前k个元素初始化到q队列中了，所以答案第一个最大就是q队头
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		//看q队列是否在窗口范围中，如果不在q删除对头
		for q[0] <= i-k {
			q = q[1:]
		}
		//已经在范围中，答案加入对头
		ans = append(ans, nums[q[0]])
	}
	return ans
}
