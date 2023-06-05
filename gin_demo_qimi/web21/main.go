package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/**
  @author: CodeWater
  @since: 2023/6/5
  @desc:
	1. gorm的默认值设置
	2. gorm的debug
**/

type User struct {
	ID int
	//这种的设置默认如果是插入null的数据，那么其实还是会是alice
	Name string `gorm:"default:'alice'"`
	Age  int64
	//默认值设置，这种修改之后，创建数据的时候会把null，空字符串、false也插入进去
	TestAddr  sql.NullString `gorm:"default:'test_null'"`
	TestPhone *int16
}

func main() {
	//args := "root:mysql's password@(ip address:port)/db1?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", args)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	//valid表示空值有效
	u := User{Name: "code", Age: 18, TestAddr: sql.NullString{String: "", Valid: true}, TestPhone: new(int16)}
	u1 := User{Name: "", Age: 18, TestAddr: sql.NullString{String: "plus", Valid: true}, TestPhone: new(int16)}
	fmt.Println(db.NewRecord(&u))
	//加个debug可以看到执行过程的语句
	db.Debug().Create(&u)
	db.Debug().Create(&u1)
	fmt.Println(db.NewRecord(&u))

}
