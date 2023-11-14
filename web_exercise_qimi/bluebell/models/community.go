package models

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}
