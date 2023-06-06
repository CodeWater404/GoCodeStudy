package models

/**
  @author: CodeWater
  @since: 2023/6/6
  @desc: 关于数据库增删改查的操作都在这个包下
	PS: 由于操作简单，就没有把controller张的操作迁移过来
**/

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
