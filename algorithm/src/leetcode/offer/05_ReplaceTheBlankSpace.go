package offer

import "strings"

/**
  @author: CodeWater
  @since: 2023/6/7
  @desc: 替换空格
**/
func replaceSpace(s string) string {
	//声明一个切片，因为append第一个参数需要切片类型
	res := []string{}
	// 使用range迭代字符串时，返回的是Unicode码点（rune）而不是字节（byte）。因此，变量value的类型是int32，表示一个Unicode码点。
	for _, value := range s {
		if value == ' ' {
			res = append(res, "%20")
		} else {
			// fmt.Printf("===>>> value type : %T\n" , value)
			//value是int32类型
			res = append(res, string(value))
		}
	}
	//使用了strings.Join函数来将字符串切片连接为一个字符串。
	return strings.Join(res, "")
}
