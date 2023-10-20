package _1_100

/**
  @author: CodeWater
  @since: 2023/10/20
  @desc: 删除有序数组中的重复项
**/
//# 思路: removeDuplicates
//1. 双指针：一个指针i遍历所有数组，一个指针k指向结果数组尾部。
//2. k指针移动条件：
//* 当i指针遍历的当前元素不等于k的前两位元素即可后移（元素可重复两次）；
//* 另外一个条件就是，k小于2的时候。
//3. k移动之前：需要把i位置所指的元素前移到k处，也就是形成结果数组，达到删除的效果。

func removeDuplicates(nums []int) int {
	k := 0
	for _, v := range nums {
		if k < 2 || (nums[k-1] != v || nums[k-2] != v) {
			nums[k] = v
			k++
		}
	}
	return k
}
