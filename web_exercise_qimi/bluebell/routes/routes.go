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

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登录用户，判断请求头中是否有 有效的jwt
		c.String(http.StatusOK, "pong")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
