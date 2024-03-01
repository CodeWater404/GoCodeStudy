package _1_80

/**
  @author: CodeWater
  @since: 2024/3/1
  @desc: 69. x 的平方根
**/

func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r + 1) >> 1
		if mid <= x/mid {
			//mid^2 <= x:那么平方根一定在mid到r之间，更新l
			l = mid
		} else {
			r = mid - 1
		}
	}
	return r
}
