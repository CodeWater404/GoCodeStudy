package offer

import "sort"

/**
  @author: CodeWater
  @since: 2023/6/16
  @desc: 在排序数组中查找数字 I
	基本思路：二分
**/
func search1(nums []int, target int) int {
	//SearchInts在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
	leftIndex := sort.SearchInts(nums, target)
	//找不到指定元素值
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return 0
	}
	// 右边界
	rightIndex := sort.SearchInts(nums, target+1) - 1
	return rightIndex - leftIndex + 1
}

func search2(nums []int, target int) int {
	n := len(nums)
	biSearch := func(condition func(int, int) bool) int {
		l, r := -1, n
		for l+1 != r {
			m := (l + r) >> 1
			if condition(m, target) {
				l = m
			} else {
				r = m
			}
		}
		return l
	}
	r := biSearch(func(m, target int) bool { return nums[m] <= target })
	l := biSearch(func(m, target int) bool { return nums[m] < target })
	return r - l
}

func search3(nums []int, target int) int {
	low, high, cnt := 0, len(nums)-1, 0
	for low <= high {
		mid := (low + high) >> 1
		if target == nums[mid] {
			for low <= high {
				if nums[low] == target {
					cnt++
				}
				low++
			}
		}
		if target < nums[mid] {
			high = mid - 1
		}
		if target > nums[mid] {
			low = mid + 1
		}
	}
	return cnt
}
