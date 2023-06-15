package offer

/**
  @author: CodeWater
  @since: 2023/6/14
  @desc: 滑动窗口的最大值
**/

//method a================================================================
func maxSlidingWindow1(nums []int, k int) []int {
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

//method b================================================================
//优先队列，简洁版
func maxSlidingWindow2(nums []int, k int) []int {
	length, hh, tt := len(nums), 0, -1
	if length == 0 {
		return []int{0}
	}
	//q优先队列：保持数值递减，但是实际存储的是数值在数组中的下标； res存储答案
	q, res := make([]int, length), make([]int, length-k+1)

	for i := 0; i < length; i++ {
		//q队列有元素并且队头不在窗口范围内，删除队头
		if hh <= tt && q[hh] < i-k+1 {
			hh++
		}
		//队列有元素并且队尾表示的元素值小于当前遍历的数组值，删除队尾
		for hh <= tt && nums[q[tt]] <= nums[i] {
			tt--
		}
		//优先队列队尾加入元素
		tt++
		q[tt] = i
		//遍历的下标超过窗口表示的范围，把res表示的下标处理一下
		if i >= k-1 {
			res[i-k+1] = nums[q[hh]]
		}
	}
	return res
}
