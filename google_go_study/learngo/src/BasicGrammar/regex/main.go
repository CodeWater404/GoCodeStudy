package main

import (
	"fmt"
	"regexp"
)

/**
  @author: CodeWater
  @since: 2023/5/22
  @desc: 正则表达式：提取邮箱的小安列
**/

const text = `
My email is codewater@mail.com@asd.com
email1 is sad@dss.com
email2 is dfss@sds.com
email3 is dsada@dsd.com.cn
`

func main() {
	//1.直接匹配
	//re := regexp.MustCompile("codewater@mail.com")
	//2.正则匹配
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
