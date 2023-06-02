package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/6/2
  @desc: gin框架中获取url、表单、前端传送的json数据（使用结构体接收）
**/

type UserInfo struct {
	//可以增加多个tag标签，然后前端请求的时候就必须要这里的tag值对应才能有正确的值
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func main() {
	r := gin.Default()

	//query-string:
	r.GET("/user", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		u := UserInfo{
			Username: username,
			Password: password,
		}

		//直接使用gin框架提供的方法直接绑定属性(后面结构有再多的属性也不需要自己处理，另外还支持表单、json数据等等)
		var u2 UserInfo
		err := c.ShouldBind(&u2)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("========>>>> user get: %v\n", u2)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})

		}

		fmt.Printf("====>>> userinfo: %v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	//获取form的测试请求，其实代码一样
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "form err",
			})
		} else {
			fmt.Printf("====>form post: %v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("====>>>> gin run failed , err: %v\n", err)
		return
	}
}
