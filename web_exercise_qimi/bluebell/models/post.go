package models

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

type Post struct {
	// 类型相同的字段可以放在一块，内存对齐，d减少内存占用
	ID          int64  `json:"id,string" db:"post_id"` //json:"id,string"在发序列化json的时候，把id转成字符。因为前端js整数类型没有64位存不下
	AuthorID    int64  `json:"author_id" db:"author_id"`
	CommunityID int64  `json:"community_id" db:"community_id" binding:"required"`
	Status      int32  `json:"status" db:"status"`
	Title       string `json:"title" db:"title" binding:"required"`
	Content     string `json:"content" db:"content" binding:"required"`
	CreateTime  string `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Post
	*CommunityDetail `json:"community"`
}
