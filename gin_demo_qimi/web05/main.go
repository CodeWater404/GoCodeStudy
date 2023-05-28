package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/5/26
  @desc: 模板语法详解

**/

type User struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//打开模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed , err: %v\n", err)
		return
	}
	//渲染模板

	u1 := User{
		Name:   "codewater",
		Age:    18,
		Gender: "male",
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"羽毛球",
	}
	m1 := map[string]interface{}{
		"Name":   "codewater2",
		"Age":    22,
		"Gender": "female",
		"hobby":  hobbyList,
	}

	//err = t.Execute(w, u1)
	err = t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
	if err != nil {
		fmt.Printf("rander template falied, err: %v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http server start failed , err: %v\n", err)
	}
}
