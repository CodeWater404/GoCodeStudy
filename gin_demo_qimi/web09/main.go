package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/5/31
  @desc: gin框架中的html渲染
	1. 对于同名模板文件的处理：定义模板名字（如果没有定于名义，那名字就是文件名）
	2. 自定义模板函数
	3，静态文件处理（html，css，js）
	4. 从网上随便下载前端的静态网站模板，然后生成运行一个网站（需要注意路径的修改，并把index.html放到自己的template下面当作home）
**/

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//gin中自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	//gin中加载statics文件(当前端请求路径是xxx的时候，就会转到statics下)
	r.Static("/xxx", "./statics")

	//加载html文件（单个模板文件的时候可以）
	//r.LoadHTMLFiles("templates/index.tmpl")
	//r.LoadHTMLFiles("templates/posts/index.tmpl" , "templates/users/index.tmpl")
	//使用正则匹配加载文件文件（对于模板文件很多的时候）
	r.LoadHTMLGlob("templates/**/*")

	//解析
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Hello , codewater!   posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Hello , users",
			"html":  "<a href='https://www.baidu.com'>baidu</a>",
		})
	})
	//返回随便从网上下载的一个网站模板
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("===========>>>> gin run failed , err: %v\n", err)
		return
	}
}
