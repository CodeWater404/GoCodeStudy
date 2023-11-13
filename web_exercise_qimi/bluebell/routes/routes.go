package routes

import (
	"net/http"
	"strings"
	"web_exercise_qimi/bluebell/controller"
	"web_exercise_qimi/bluebell/logger"
	"web_exercise_qimi/bluebell/pkg/jwt"

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
	r.GET("/ping", JWTAuthMiddleware(), func(c *gin.Context) {
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

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头中 2.放在请求体中 3.放在URI中
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization：Bearer xxx.xxx.xxx / X-TOKEN: xxx.xxx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("userID", mc.UserID)

		c.Next() // 后续的处理函数可以用过c.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}
