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
func SignUp(p *models.ParamSignUp) {
	// 1. 判断用户是否存在
	mysql.QueryUserByUsername()
	// 2. 生成UID
	snowflake.GenID()
	// 3. 存入数据库
	mysql.InsertUser()
}
