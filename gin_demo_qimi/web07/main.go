package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/5/30
  @desc: 模板继承
**/

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("index=====>>>>template parse files failed , err:%s\n", err)
		return
	}
	msg := "codewater"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("index====>>>>>> execute failed , err: %s\n", err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("home=====>>>>parse files failed , err: %s\n", err)
		return
	}
	msg := "codewater"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("home=====>>>> execute failed , err: %s\n", err)
		return
	}
}

func index2(w http.ResponseWriter, r *http.Request) {
	//先写夫模板,再写子模版
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("===========>>>> index2 parse falied , err: %v\n", err)
		return
	}
	name := "codewater index2"
	//因为解析了多个模板,所以这里需要指定执行模板
	err = t.ExecuteTemplate(w, "index2.tmpl", name)
	if err != nil {
		fmt.Printf("s======>>>>>> index2 execute failed , err: %v\n", err)
		return
	}
}

func home2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("===========>>>> home2 parse falied , err: %v\n", err)
		return
	}
	name := "codewater home2"
	err = t.ExecuteTemplate(w, "home2.tmpl", name)
	if err != nil {
		fmt.Printf("s======>>>>>> home2 execute failed , err: %v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("=====>>>http server start failed , err:%s\n", err)
	}
}
