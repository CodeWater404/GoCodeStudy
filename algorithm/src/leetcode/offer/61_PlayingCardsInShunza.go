package offer

import "sort"

/**
  @author: CodeWater
  @since: 2023/6/26
  @desc: 扑克牌中的顺子
**/
func isStraight(nums []int) bool {
	joker := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i+1] == nums[i] {
			//重复元素直接提前返回，不是顺序了
			return false
		}
	}
	return nums[4]-nums[joker] < 5 // 最大牌 - 最小牌的差值 < 5 则可构成顺子
}
