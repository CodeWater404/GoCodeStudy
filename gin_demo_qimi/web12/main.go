package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/6/2
  @desc: 获取表单数据
**/

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		user2 := c.DefaultPostForm("user2", "user2222, failed")
		pass2 := c.DefaultQuery("pass2", "pass2222 , failed")

		user3, ok := c.GetPostForm("user3")
		if !ok {
			user3 = "user3 failed , user3333"
		}
		pass3, err := c.GetPostForm("pass3 failed , pass333")
		if !err {
			pass3 = "pass3 failed , pass3333"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
			"user2":    user2,
			"pass2":    pass2,
			"user3":    user3,
			"pass3":    pass3,
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("=======>>> gin run failed , err:%v\n", err)
		return
	}
}
