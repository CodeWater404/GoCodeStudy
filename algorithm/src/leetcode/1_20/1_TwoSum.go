package __20

/**
  @author: CodeWater
  @since: 2023/5/14
  @desc: 两数之和
**/
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	hash := make(map[int]int)
	//the other way to initialize : hash := map[int]int
	for index, value := range nums {
		if diff, err := hash[target-value]; err {
			return []int{index, diff}
		}
		hash[value] = index
	}
	return nil
}
