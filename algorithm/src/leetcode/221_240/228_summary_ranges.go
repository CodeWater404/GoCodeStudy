package _21_240

import "strconv"

/**
  @author: CodeWater
  @since: 2024/1/3
  @desc: 228. 汇总区间
**/

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		j := i + 1 // i区间左端点，j区间右端点
		for j < len(nums) && nums[j] == nums[j-1]+1 {
			j++
		}
		if j == i+1 { //单个区间
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res = append(res, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j-1]))
		}
		//更新左端点,因为for循环会自动+1，所以这里要-1
		i = j - 1
	}
	return res
}
