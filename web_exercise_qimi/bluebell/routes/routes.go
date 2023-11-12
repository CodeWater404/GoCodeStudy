package routes

import (
	"net/http"
	"web_exercise_qimi/bluebell/controller"
	"web_exercise_qimi/bluebell/logger"

	"github.com/gin-gonic/gin"
)

/*
*

	@author: CodeWater
	@since: 2023/11/11
	@desc: $

*
*/
func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
