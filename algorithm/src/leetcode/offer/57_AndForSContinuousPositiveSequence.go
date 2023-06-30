package offer

/**
  @author: CodeWater
  @since: 2023/6/30
  @desc: 和为s的连续正数序列
**/
func findContinuousSequence(target int) [][]int {
	//双指针：i左边界从1开始，j右边界从2开始，s当前i到j范围内的和
	i, j, s := 1, 2, 3
	res := [][]int{}
	for i < j {
		//找到一组解
		if s == target {
			ans := make([]int, j-i+1)
			//从i到j
			for k := i; k <= j; k++ {
				ans[k-i] = k
			}
			res = append(res, ans)
		}
		//s和target之间有三种关系：大于、小于、等于（大于和等于可以都是一样的操作放一起）
		if s >= target {
			//s大于等于target了，移动左边界减小s
			s -= i
			i++
		} else {
			//s小于target，移动右边界增加s的大小
			j++
			s += j
		}
	}
	return res
}
