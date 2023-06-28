package offer

/**
  @author: CodeWater
  @since: 2023/6/28
  @desc: 不用加减乘除做加法
**/
func add(a int, b int) int {
	for b != 0 {
		//c进位
		c := (a & b) << 1
		//a非进位和
		a ^= b
		b = c
	}
	return a
}
