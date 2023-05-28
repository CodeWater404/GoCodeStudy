package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/5/28
  @desc: 模板嵌套
**/

/**f
** @Description: 自定义函数传给模板文件，需要注意的是在解析模板文件之间先声明函数，Funcs
** @param w
** @param r
**/
func f(w http.ResponseWriter, r *http.Request) {
	praise := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	t := template.New("f.tmpl")
	t.Funcs(template.FuncMap{
		"kua": praise,
	})
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse templates files failed , err: %s\n", err)
		return
	}
	name := "codewater"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("execute template failed , err: %s\n", err)
	}

}

func demo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("======>>>>demo func parse templates failed , err:%s\n", err)
		return
	}

	name := "codewater"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("====>>>demo func execute failed , err: %s\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", f)
	http.HandleFunc("/tmpl", demo)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http server start failed , err: %s\n", err)
		return
	}
}
