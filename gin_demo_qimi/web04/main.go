package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/5/23
  @desc: go原生——模板文件渲染
**/

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板:当前目录下的hello.tmpl.(注意这里用命令去创建exe去运行，因为如果是IDE的话，很可能生成的exe不在当前目录下，那么运行的时候就会找不到这个模板文件)
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse temolate failed , err: %v", err)
		return
	}
	//渲染模板
	name := "codewater"
	err = t.Execute(w, name)

	if err != nil {
		fmt.Printf("render template failed , err: %v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed , err：%v", err)
	}

}
