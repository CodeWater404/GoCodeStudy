package models

/**
  @author: CodeWater
  @since: 2023/11/12
  @desc: $
**/

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
