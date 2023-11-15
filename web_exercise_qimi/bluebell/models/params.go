package models

/**
  @author: CodeWater
  @since: 2023/11/12
  @desc: 定义参数的结构体
**/

// ParamSignUp 用于接收注册请求参数的结构体
type ParamSignUp struct {
	Username   string `json:"username"  binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //eqfield=Password 与 Password 字段相等; required 必填
}

// ParamLogin 用于接收登录请求参数的结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	// UserID 从请求中获取当前用户的ID
	PostID    string `json:"post_id" binding:"required"`       // 帖子ID
	Direction int8   `json:"direction" binding:"oneof=-1 1 0"` // 赞成票(1)还是反对票(-1)，0取消(前端传0的时候binding会过滤掉，然后required又检测不到值，所以去掉)
}
