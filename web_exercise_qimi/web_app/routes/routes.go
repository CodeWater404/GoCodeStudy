package routes

import (
	"net/http"
	"web_exercise_qimi/web_app/logger"

	"github.com/gin-gonic/gin"
)

/*
*

	@author: CodeWater
	@since: 2023/11/11
	@desc: $

*
*/
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}
