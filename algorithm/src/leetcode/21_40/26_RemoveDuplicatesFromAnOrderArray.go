package _1_40

/**
  @author: CodeWater
  @since: 2023/10/19
  @desc: 删除有序数组中的重复项
**/
/*
思路：双指针。一个指针i遍历所有元素，当前元素和上一个元素相等的时候说明该元素是重复的；一个指针k遍历所有不重复的元素，只有当i指针遍历到当前元素和当前的上一个元素不等时，k才往后移同时赋值给k当前的元素。
最终k所指的就是所有不等的元素。
*/
func removeDuplicates(nums []int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		//第一个元素特判
		if i == 0 || nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
