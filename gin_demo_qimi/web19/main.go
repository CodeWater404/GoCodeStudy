package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
  @author: CodeWater
  @since: 2023/6/5
  @desc: gorm初识
**/

//数据表
type UserInfo struct {
	ID     int
	Name   string
	Gender string
	Hobby  string
}

func main() {
	args := "root:mysql's password@(ip address:port)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", args)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表  自动把结构体和表进行对应
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	//u1 := UserInfo{1, "code", "male", "蛙泳"}
	//db.Create(&u1)

	var u UserInfo
	db.First(&u) //查询表中第一条数据，保存到u中
	fmt.Printf("u: %#v\n", u)

	//更新
	db.Model(&u).Update("hobby", "双色球")

	//删除
	db.Delete(&u)
}
