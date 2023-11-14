package logic

import (
	"web_exercise_qimi/bluebell/dao/mysql"
	"web_exercise_qimi/bluebell/models"
	"web_exercise_qimi/bluebell/pkg/snowflake"

	"go.uber.org/zap"
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

func GetPostById(pid int64) (data *models.PostDetail, err error) {
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityDetailById(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}
	data = &models.PostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}
