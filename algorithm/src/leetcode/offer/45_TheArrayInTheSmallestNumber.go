package offer

import (
	"fmt"
	"sort"
)

/**
  @author: CodeWater
  @since: 2023/6/24
  @desc: 把数组排成最小的数
**/
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		//将nums[i]和nums[j]转换为字符串并拼接起来，分别赋值给变量x和y。
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		//比较拼接后的xy，这里小于实际上就是升序，也就是把拼接后的unicode码值小元素的排在前面，注意不是拼接后的xy。比如：3和30，x=330  y=303 ，x大于y，所以在nums数组中30会在3的前面
		return x < y
	})

	res := ""
	for i := 0; i < len(nums); i++ {
		res += fmt.Sprintf("%d", nums[i])
	}
	return res
}