package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/**
  @author: CodeWater
  @since: 2023/6/6
  @desc: gorm里面的一些常见的查询操作
**/

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	arg := "root:mysql's password@(ip address:port)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", arg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	//创建数据
	u1 := User{Name: "code", Age: 12}
	db.Create(&u1)
	u2 := User{Name: "jinzhu", Age: 22}
	db.Create(&u2)

	//查询，first是找到一条
	/*根据主键查询第一条记录
	db.First(&user)
	SELECT * FROM users ORDER BY id LIMIT 1;
	*/
	var user User
	db.First(&user)
	//另外一种写法，指针类型声明
	//var user new(User)
	//db.First(user)
	fmt.Printf("1111====>>>user: %#v\n", user)

	//find是找出所有
	var users []User
	db.Find(&users)
	fmt.Printf("2222======>>>users: %#v\n", users)

	//如果找不到，就用给定的参数赋给变量,但是实际数据库中不会插入这条数据
	var user3 User
	db.FirstOrInit(&user3, User{Name: "hhhhh"})
	fmt.Printf("3333======>>>user3: %#v\n", user3)

	//attrs也是给变量赋值的
	var user4 User
	db.Attrs(User{Age: 90}).FirstOrInit(&user4, User{Name: "hhhhhhhh"})
	fmt.Printf("44444======>>>>user4: %#v\n", user4)

	//FirstOrCreate:获取匹配的第一条记录, 否则根据给定的条件创建一个新的记录 (仅支持 struct 和 map 条件)
}
