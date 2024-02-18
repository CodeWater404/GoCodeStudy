package _01_220

/**
  @author: CodeWater
  @since: 2024/2/18
  @desc: 215. 数组中的第K个最大元素
**/

func findKthLargest(nums []int, k int) int {
	return quick_sort(nums, 0, len(nums)-1, k-1)
}

// quick_sort 降序
func quick_sort(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	x, i, j := nums[l], l-1, r+1
	for i < j {
		//降序，所以这里是nums[i] > x
		for i++; nums[i] > x; i++ {
		}
		for j--; nums[j] < x; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j { // k<=j，说明在l-j这段区间里面
		return quick_sort(nums, l, j, k)
	} else {
		return quick_sort(nums, j+1, r, k)
	}
}
