// Code generated by goctl. DO NOT EDIT.
package types

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoRequest struct {
	ID int `path:"id"`
}

type UserInfoResponse struct {
	UserId    int    `json:"user_id"`
	UserrName string `json:"username"`
}