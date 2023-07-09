package _61_180

/**
  @author: CodeWater
  @since: 2023/7/8
  @desc: 两数之和——输入有序的数组
**/
func twoSum(numbers []int, target int) []int {
	for i , j := 0 , len(numbers) - 1 ; i < j ; i++ {
		for i < j && numbers[i] + numbers[j] > target {
			j--
		}
		if i < j && numbers[i] + numbers[j] == target {
			//+1是因为下标从1开始
			return []int{i + 1 , j + 1}
		}
	}
	//因为题目保证有唯一的解
	return []int{}
}