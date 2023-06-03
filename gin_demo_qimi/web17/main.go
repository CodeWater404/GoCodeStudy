package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/6/3
  @desc: 路由
	1. 处理不存在的路由
	2. 匹配任何类型的路由
	3. 路由组（url前缀一致的）
**/

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "post",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "put",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})
	})

	//可以接收一切请求类型的
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "any get"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "any post"})
		}
		c.JSON(http.StatusOK, gin.H{
			"method": "any",
		})
	})

	//如果用户访问到不存在的错误，处理这个请求
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Not found......codewater",
		})
	})

	//路由组({}还必须换行写。。。。)
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/xx",
			})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/oo",
			})
		})
	}

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("=====>>>> gin run failed , err: %v\n", err)
		return
	}
}
