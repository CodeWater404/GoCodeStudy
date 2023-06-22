package offer

/**
  @author: CodeWater
  @since: 2023/6/22
  @desc: 求1+2+3+。。+n
**/
func sumNums(n int) int {
	var res int = 0
	//因为题目要求：不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句
	var helper func(a int) bool
	helper = func(a int) bool {
		res += a
		//这里使用&&来实现递归的终止退出条件
		return a > 1 && helper(a-1)
	}
	helper(n)
	return res
}
