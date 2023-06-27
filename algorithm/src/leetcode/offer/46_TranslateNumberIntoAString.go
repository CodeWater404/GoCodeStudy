package offer

import "strconv"

/**
  @author: CodeWater
  @since: 2023/6/27
  @desc: 把数字翻译成字符串
**/
func translateNum(num int) int {
	//转换为字符串
	s := strconv.Itoa(num)
	//a:0个字符时的翻译方法数（由2个字符-1个字符推算出来）；b：单个字符时的翻译方法数
	a, b := 1, 1
	for i := 2; i <= len(s); i++ {
		//尝试将两位字符组合起来
		tmp := s[i-2 : i]
		c := a
		if tmp >= "10" && tmp <= "25" {
			//两位字符可以翻译，将前两种状态方案数加起来等于前i个字符的方案书
			c += b
		}
		//状态顺移：b记录前i-2种字符的方案数，a记录前i-1种方案数
		b = a
		a = c
	}

	return a
}
