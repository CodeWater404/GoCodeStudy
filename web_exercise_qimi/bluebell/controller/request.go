package controller

import (
	"errors"
	"web_exercise_qimi/bluebell/middleware"

	"github.com/gin-gonic/gin"
)

/**
  @author: CodeWater
  @since: 2023/11/13
  @desc: $
**/

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登录的用户ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middleware.CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
