// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type UserInfo struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"username"`
}

type UserInfoResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}
