package offer

/**
  @author: CodeWater
  @since: 2023/6/29
  @desc: 数组中出现次数超过一半的数字
**/
func majorityElement(nums []int) int {
	//摩尔投票法：假设x为众数，votes为票数.由于众数超过一般，所以众数和非众数互相抵消票数之后，只剩下众数x
	x, votes := 0, 0
	for _, num := range nums {
		//votes票数为0，设当前数为众数
		if votes == 0 {
			x = num
		}
		//如果当前数num与众数x相等，票数+1；不等票数-1
		if num == x {
			num = 1
		} else {
			num = -1
		}
		votes += num
	}

	return x
}
