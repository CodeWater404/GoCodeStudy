package _81_300

/**
  @author: CodeWater
  @since: 2024/3/7
  @desc: 300. 最长递增子序列
**/

func lengthOfLIS(nums []int) int {
	q := make([]int, 0)
	for _, x := range nums {
		if len(q) == 0 || x > q[len(q)-1] {
			q = append(q, x)
		} else {
			if x <= q[0] {
				q[0] = x
			} else {
				l, r := 0, len(q)-1
				for l < r {
					mid := (l + r + 1) >> 1
					if q[mid] < x {
						l = mid
					} else {
						r = mid - 1
					}
				}
				q[r+1] = x
			}
		}
	}
	return len(q)
}
