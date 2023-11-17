package controller

import (
	"strconv"
	"web_exercise_qimi/bluebell/logic"
	"web_exercise_qimi/bluebell/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

// CreatePostHandler 创建帖子
// @Summary 创建帖子
// @Description 创建帖子,需要登录,会放入数据库和redis中
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /api/v1/post [post]
func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreatePost with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 获取当前发帖的用户ID
	userID, err := getCurrentUser(c) //todo:用postman测试时，auth是jwt bear userid=0，是bear token userid可以获取到
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

// GetPostDetailHandler 获取帖子详情
// @Summary 获取帖子详情
// @Description 获取帖子详情，会展示帖子内容
// @Tags 帖子相关接口
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id query int true "查询参数，帖子id"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /api/v1/post/{id} [get]
func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表
// @Summary 获取帖子列表
// @Description 获取帖子列表,会展示帖子内容
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /api/v1/posts [get]
func GetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler2 获取帖子列表[新]
// @Summary 获取帖子列表[新]
// @Description 获取帖子列表,会展示帖子内容,可以按照时间或者分数排序,默认按照时间排序,可以指定社区,不指定社区就是查所有的帖子
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /api/v1/posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// get请求参数（query string）： /api/v1/posts2/?page=1&size=10&order=time
	// 初始化结构体时指定初始参数，前端没传的时候就用这个当作默认值
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, //用常量来代替字符串，避免写错
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p)
	if err != nil {
		zap.L().Error("logic.GetPostListNew failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
