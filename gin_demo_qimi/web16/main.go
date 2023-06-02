package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/6/2
  @desc: 重定向
**/

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "ok",
		//})

		//这种重定向，地址栏会改变
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/a", func(c *gin.Context) {
		//这种重定向,地址栏不会改变
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("=====>>>> gin run failed , err:%v\n", err)
		return
	}
}
