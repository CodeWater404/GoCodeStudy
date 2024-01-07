package init_gorm

import (
	"fmt"
	"go-zero/feng/model_study/user_gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
  @author: CodeWater
  @since: 2024/1/7
  @desc: $
**/

func InitGorm(MysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{})
	if err != nil {
		panic("connect to mysql failed , error:" + err.Error())
	}
	fmt.Println("connect to mysql success!!!!")
	db.AutoMigrate(&models.UserModel{})

	return db
}
