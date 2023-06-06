package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/**
  @author: CodeWater
  @since: 2023/6/6
  @desc: 专门连接数据库操作的包
**/

var (
	DB *gorm.DB
)

/**initMysql
** @Description: mysql初始化
** @return err
**/
func InitMysql() (err error) {
	arg := "root:mysql's password@(ip address:port)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", arg)
	if err != nil {
		fmt.Printf("======>>> gorm connect failed , err:%v\n", err)
		return
	}
	return DB.DB().Ping()
}

func Close() {
	defer DB.Close()
}
