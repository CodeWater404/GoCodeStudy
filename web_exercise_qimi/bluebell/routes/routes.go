package routes

import (
	"net/http"
	"web_exercise_qimi/bluebell/controller"
	"web_exercise_qimi/bluebell/logger"
	"web_exercise_qimi/bluebell/middleware"

	"github.com/gin-gonic/gin"
)

/*
*

	@author: CodeWater
	@since: 2023/11/11
	@desc: $

*
*/

// Setup 路由的设置
func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	//todo:写成这样air报错乱码v1.Use(middlelware.JWTAuthMiddleware())
	v1.Use(middleware.JWTAuthMiddleware()) //这里明明写对了，但是goland还是报错，不知道为什么（Unresolved reference 'JWTAuthMiddleware'）

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)

		v1.POST("/vote", controller.PostVoteHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
