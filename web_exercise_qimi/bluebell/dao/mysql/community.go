package mysql

import (
	"database/sql"
	"web_exercise_qimi/bluebell/models"

	"go.uber.org/zap"
)

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id , community_name from community"
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no community in the table community")
			err = nil
		}
	}
	return
}
