package _1_40

/**
  @author: CodeWater
  @since: 2023/10/19
  @desc:移除元素
**/
/*
思路：双指针，一个i遍历数组；一个k只有当另外一个指针遍历到不等于val的时候才往后移，也就是实现把等于val的被覆盖掉了。最终，k的值就是移除后的数组长度
*/
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
