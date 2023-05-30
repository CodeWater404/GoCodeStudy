package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/5/30
  @desc: 模板自定义标识符
**/

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("======>>>> index parse faild , err: %v\n", err)
		return
	}
	name := "codewater"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("=====>>>> index execute faild , err: %v\n", err)
		return
	}
}

/**xss
** @Description: html下的template会把一些前端的东西转义掉,这样用户就不能恶意攻击
** @param w
** @param r
**/
func xss(w http.ResponseWriter, r *http.Request) {
	//通过实现自定义函数传给模板,实现对有些内容的控制,不需要转义
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("=====>>>> xxx parse faild , err:%v\n", err)
		return
	}
	//这里模拟用户的恶意攻击(如果使用的是html下的template,那么这里就会被转移掉,如果是text下的,那么就不会转义)
	str1 := "<script>alert(123)</script>"
	str2 := "<a href='http://www.baidu.com'>baidu</a>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Printf("======>>>>>xss execute faild , err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("=====>>>>main listen faild , err: %v\n", err)
		return
	}
}
