package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

/**
  @author: CodeWater
  @since: 2023/5/11
  @desc: $
**/

func main() {
	//创建一个服务
	ginServer := gin.Default()
	//设置网页标签上的图标
	ginServer.Use(favicon.New("./favicon.ico"))
	//访问地址，处理我们的请求 Request Response
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})
	//restful api
	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "post user"})
	})
	ginServer.PUT("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "put user"})
	})
	ginServer.DELETE("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "delete user"})
	})
	//服务器端口
	ginServer.Run(":8082")
}
