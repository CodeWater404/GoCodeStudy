package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/**
  @author: CodeWater
  @since: 2023/6/6
  @desc: gorm的删除操作
	1. 有deleted_at字段，是乱删除；如果没有这个字段那么删除的时候就会直接物理删除
	2. 软删除的话也可以调用方法物理删除
**/

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	arg := "root:mysql's password@(ip address:port)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", arg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	u1 := User{Name: "code", Age: 12, Active: true}
	db.Create(&u1)
	u2 := User{Name: "kody", Age: 22, Active: false}
	db.Create(&u2)

	// delete
	var u User
	//软删除，必需是主键才会删除指定的数据；如果不是主键，那么会删除该数据表上的所有数据！！！！！！！！
	u.ID = 1
	db.Debug().Delete(&u)
	//删除指定数据，这个不会删除所有数据，是因为先用where查询的
	db.Debug().Where("name=?", "kody").Delete(User{})
	db.Debug().Delete(User{}, "age=?", 12)

	u3 := User{Name: "code2", Age: 33, Active: true}
	db.Create(&u3)
	u4 := User{Name: "kody2", Age: 44, Active: false}
	db.Create(&u4)
	u6 := User{Name: "kody2", Age: 55, Active: false}
	db.Create(&u6)

	var u5 []User
	db.Debug().Where("name=?", "kody2").Find(&u1)
	fmt.Printf("=====>>>u5: %v\n", u5)

	//物理删除Unscoped()
	db.Debug().Unscoped().Where("name=?", "code2").Delete(User{})
}
