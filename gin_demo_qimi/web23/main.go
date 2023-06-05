package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/**
  @author: CodeWater
  @since: 2023/6/6
  @desc: gorm update operations
	1. save：更新所有
	2. update：更新一列
	3. updates： 更新多个列
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

	//把结构体和数据表相对应
	db.AutoMigrate(&User{})

	u1 := User{Name: "code", Age: 19, Active: true}
	db.Create(&u1)
	u2 := User{Name: "cidy", Age: 22, Active: false}
	db.Create(&u2)

	var user User
	db.First(&user)

	//update
	user.Name = "ccccc"
	user.Age = 88
	db.Debug().Save(&user) // update all columns
	//更新指定字段，但是一些默认字段也会更新
	/*例如这些字段：
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	*/
	db.Debug().Model(&user).Update("name", "c22222")
	//UPDATE `users` SET `name` = 'c22222', `updated_at` = '2023-06-06 02:27:31'  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1

	m1 := map[string]interface{}{
		"name":   "c3333",
		"age":    33,
		"active": true,
	}
	//更新m1列出来的所有字段
	db.Debug().Model(&user).Updates(m1)
	//只更新指定列
	db.Debug().Model(&user).Select("age").Update(m1)
	//除了active其他列都更新
	db.Debug().Model(&user).Omit("active").Updates(m1)
	//更新指定字段，但是不会更新hooks，比如：UpdatedAt
	db.Debug().Model(&user).UpdateColumn("age", 44)

	rowsNum := db.Debug().Model(User{}).Updates(User{Name: "hello", Age: 55}).RowsAffected
	fmt.Printf("====>>>>rowsNum: %v\n", rowsNum)

	//让user表中的所有用户年龄加2
	db.Debug().Model(&User{}).Update("age", gorm.Expr("age + ?", 2))

}
