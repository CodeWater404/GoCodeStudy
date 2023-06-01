package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/6/1
  @desc: gin框架
	1. 返回json数据,
	2, 几种写法
**/

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//write method one
		//data := map[string]interface{}{
		//	"name":    "codewater",
		//	"message": "hello , json",
		//	"status":  "200",
		//}

		//write method two
		data := gin.H{"name": "code", "message": "hello , data2", "status": "ok"}
		c.JSON(http.StatusOK, data)
	})

	type msg struct {
		//Name大写是可访问，但是如果传给前端或者别的什么地方需要小写，那么可以加上`要用到包的包名：”属性名“`
		Name    string `json:"name"`
		Message string
		Status  string
	}
	r.GET("/json2", func(c *gin.Context) {
		data := msg{
			"code",
			"hello , struct message",
			"ok",
		}
		c.JSON(http.StatusOK, data)
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("=======>>>> gin run falied , err:%v\n", err)
		return
	}
}
