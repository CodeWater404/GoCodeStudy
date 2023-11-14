package mysql

import "errors"

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
)
