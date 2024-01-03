package _01_220

/**
  @author: CodeWater
  @since: 2024/1/3
  @desc: 219. 存在重复元素 II
**/

func containsNearbyDuplicate(nums []int, target int) bool {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if _, ok := hash[x]; ok && i-hash[x] <= target {
			return true
		}
		//如果已经存在，但是两下标超过target也需要到这更新
		hash[x] = i
	}
	return false
}
