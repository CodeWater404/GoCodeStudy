package _61_180

/*
*

	@author: CodeWater
	@since: 2023/7/8
	@desc: 两数之和——输入有序的数组

*
*/

// twoSum 双指针，acwing
func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; i++ {
		for i < j && numbers[i]+numbers[j] > target {
			j--
		}
		if i < j && numbers[i]+numbers[j] == target {
			//+1是因为下标从1开始
			return []int{i + 1, j + 1}
		}
	}
	//因为题目保证有唯一的解
	return []int{}
}

// twoSum2 双指针,自己的写法
func twoSum2(numbers []int, target int) []int {
	res := make([]int, 2)
	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			res[0], res[1] = i+1, j+1
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return res
}
