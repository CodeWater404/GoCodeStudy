package _81_200

/**
  @author: CodeWater
  @since: 2023/10/21
  @desc: 轮转数组
**/
//方法一：开辟额外空间存储结果数组
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

/*
原地做法（三次翻转）：首先将整个数组反转一遍，这样整个数组逆序，同样位于末尾k个位置的就到了数组前面；但是还是需要把第一次反转后的前k个位置的数再反转一边，k长度之后的后半段的再反转一遍，这样所有的元素就达成顺序了，也就得到结果数组。例子：
输入：[1,2,3,4,5,6,7] and k = 3
输出：[5,6,7,1,2,3,4]
解释：
移动一次后的结果: [7,1,2,3,4,5,6]
移动两次后的结果: [6,7,1,2,3,4,5]
移动三次后的结果: [5,6,7,1,2,3,4]
*/
func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	// 整个翻转
	reverse(nums)
	// 前半段翻转
	reverse(nums[0:k])
	// 后半段翻转
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
