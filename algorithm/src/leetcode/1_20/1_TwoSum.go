package __20

/**
  @author: CodeWater
  @since: 2023/5/14
  @desc: 两数之和
**/

/**twoSum1
** @Description: the enumeration of violence
** @param nums
** @param target
** @return []int
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

/**twoSum2
** @Description: hash. The difference is the key , and the subscript of the difference is the value
** @param nums
** @param target
** @return []int
**/
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
