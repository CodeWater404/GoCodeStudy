package _1_55

/**
  @author: CodeWater
  @since: 2023/10/23
  @desc:跳跃游戏
**/
func canJump(nums []int) bool {
	// i遍历整个数组；j作为能够跳到的位置
	for i, j := 0, 0; i < len(nums); i++ {
		// 如果j小于i，说明无法跳到
		if j < i {
			return false
		}
		// 否则更新能跳到的j位置，用i位置加i位置能够跳的长度
		j = max(j, i+nums[i])
	}
	// 最终没有false的话就是能够跳到数组末尾
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
