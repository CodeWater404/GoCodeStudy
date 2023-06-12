package offer

/**
  @author: CodeWater
  @since: 2023/6/12
  @desc:调整数组顺序使奇数位于偶数前面
**/

func exchange(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		for nums[i]%2 != 0 && i < j {
			i++
		}
		for nums[j]%2 == 0 && i < j {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		//交换后也要更新一下下标，不然数组偶数个在中间位置的时候又会换回去
		i++
		j--
	}
	return nums
}
