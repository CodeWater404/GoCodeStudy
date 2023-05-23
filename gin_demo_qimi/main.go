package main

import "github.com/gin-gonic/gin"

/**
  @author: CodeWater
  @since: 2023/5/15
  @desc: gin框架的入门使用
**/

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang!",
	})
}

func main() {
	//返回默认的路由引擎
	r := gin.Default()

	r.GET("/hello", sayHello)

	//some simple example for restful
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})

	r.Run(":9090")
}
