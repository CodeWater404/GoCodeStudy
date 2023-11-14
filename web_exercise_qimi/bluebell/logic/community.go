package logic

import (
	"web_exercise_qimi/bluebell/dao/mysql"
	"web_exercise_qimi/bluebell/models"
)

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
