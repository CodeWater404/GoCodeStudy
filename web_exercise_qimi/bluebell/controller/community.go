package controller

import (
	"web_exercise_qimi/bluebell/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

// CommunityHandler 社区列表
func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不能轻易暴露服务器错误
		return
	}
	ResponseSuccess(c, data)
}
