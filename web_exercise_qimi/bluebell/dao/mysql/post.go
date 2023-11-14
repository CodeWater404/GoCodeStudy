package mysql

import "web_exercise_qimi/bluebell/models"

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id , title , content , author_id , community_id) 
				values(? , ? , ? , ? , ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}
