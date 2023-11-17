package routes

import (
	"net/http"
	"time"
	"web_exercise_qimi/bluebell/controller"
	"web_exercise_qimi/bluebell/logger"
	"web_exercise_qimi/bluebell/middleware"

	/*todo:为什么这个不行
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	*/
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	_ "web_exercise_qimi/bluebell/docs" // 千万不要忘了导入把你上一步生成的docs

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
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(2*time.Second, 1)) // 令牌桶限流,2秒放一个令牌，桶子的容量为1

	v1 := r.Group("/api/v1")

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 注册业务路由r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
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
		// 根据时间或者分数获取帖子列表
		v1.GET("/posts2", controller.GetPostListHandler2)

		v1.POST("/vote", controller.PostVoteHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
