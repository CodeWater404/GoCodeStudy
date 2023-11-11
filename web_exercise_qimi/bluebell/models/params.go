package models

/**
  @author: CodeWater
  @since: 2023/11/12
  @desc: 定义参数的结构体
**/

// ParamSignUp 用于接收注册请求参数的结构体
type ParamSignUp struct {
	Username   string `json:"username" `
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
