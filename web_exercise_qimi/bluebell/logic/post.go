package logic

import (
	"web_exercise_qimi/bluebell/dao/mysql"
	"web_exercise_qimi/bluebell/models"
	"web_exercise_qimi/bluebell/pkg/snowflake"
)

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

func CreatePost(p *models.Post) error {
	p.ID = snowflake.GenID()
	return mysql.CreatePost(p)
}

func GetPostById(pid int64) (*models.Post, error) {
	return mysql.GetPostById(pid)
}
