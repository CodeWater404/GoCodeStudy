package offer

import "strings"

/**
  @author: CodeWater
  @since: 2023/6/13
  @desc: $
**/

/**reverseWords
** @Description: methods a
** @param s
** @return string
**/
func reverseWords(s string) string {
	//fields将字符串按空格分割为字符串切片，都不用trim去除空格了
	str := strings.Fields(s)
	//string不可变，用builder
	res := strings.Builder{}
	//str[0]直接就是一个单词了.  the   []string strings.Builder
	// fmt.Printf("%v %T %T\n " , str[0] , str , res)
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == " " {
			continue
		}
		//追加字符串
		res.WriteString(str[i] + " ")
	}
	return strings.TrimSpace(res.String())
}

/**reverseWords2
** @Description:doublePointer
** @param s
** @return string
**/
func reverseWords2(s string) string {
	var res string
	i := len(s) - 1
	j := i
	//双指针：i遍历到单词的首部， j每次更新到单词的尾部
	for i >= 0 {
		// 去除空格
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
		//找到单词的首部
		for i >= 0 && s[i] != ' ' {
			i--
		}
		//i+1是因为退出上面for的时候i处在空格下标上；j+1是因为切片左闭右开
		res = res + s[i+1:j+1] + " "
	}
	return strings.TrimSpace(res)
}
