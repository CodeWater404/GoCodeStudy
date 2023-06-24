package offer

/**
  @author: CodeWater
  @since: 2023/6/24
  @desc: 数组中的逆序对
**/
var temp []int

func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	temp = make([]int, len(nums))
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) int {
	if l >= r {
		return 0
	}
	//注意这里要加上（），不然会先运算r右移一位
	mid := (l + r) >> 1
	res := mergeSort(nums, l, mid) + mergeSort(nums, mid+1, r)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			res += mid - i + 1
			temp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = nums[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = nums[j]
		k++
		j++
	}
	for i, j = l, 0; i <= r; i, j = i+1, j+1 {
		nums[i] = temp[j]
	}
	return res
}
