package __20

/*
*

	@author: CodeWater
	@since: 2023/12/27
	@desc: 11. 盛最多水的容器

*
*/
func maxArea(height []int) int {
	res := 0
	for i, j := 0, len(height)-1; i < j; {
		res = max(res, (j-i)*min(height[i], height[j]))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
