package logic

import (
	"web_exercise_qimi/bluebell/dao/mysql"
	"web_exercise_qimi/bluebell/models"
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
