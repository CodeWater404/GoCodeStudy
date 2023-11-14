package logic

import (
	"web_exercise_qimi/bluebell/dao/mysql"
	"web_exercise_qimi/bluebell/models"
	"web_exercise_qimi/bluebell/pkg/jwt"
	"web_exercise_qimi/bluebell/pkg/snowflake"
)

/*
*

	@author: CodeWater
	@since: 2023/11/12
	@desc: $

*
*/

// SignUp 处理注册的业务逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2. 生成UID
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3. 存入数据库
	return mysql.InsertUser(user)
}

// Login 处理登录的业务逻辑
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
