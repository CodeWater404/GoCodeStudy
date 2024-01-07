package models

import "gorm.io/gorm"

/**
  @author: CodeWater
  @since: 2024/1/7
  @desc: $
**/

type UserModel struct {
	gorm.Model
	Username string `gorm:"size:32" json:"username"`
	Password string `gorm:"size64" json:"password"`
}
