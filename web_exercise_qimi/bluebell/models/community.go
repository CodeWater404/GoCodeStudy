package models

import "time"

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"` //omitempty 如果为空则不返回
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
