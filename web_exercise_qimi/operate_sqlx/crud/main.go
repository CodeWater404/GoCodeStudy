package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/11/7
  @desc: crud with sqlx
**/

var (
	db  *sqlx.DB
	dsn = "root:%s@tcp(%s:3306)/gin_qimi?charset=utf8mb4&parseTime=True"
)

func initDB() (err error) {
	ten, tenDB := os.Getenv("tencent"), os.Getenv("tencent_mysql")
	if ten == "" || tenDB == "" {
		err = errors.New(fmt.Sprintf("get env failed , ten:%s , tenDb:%s\n", ten, tenDB))
		return
	}

	dsn = fmt.Sprintf(dsn, tenDB, ten)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect db failed , err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	//这里sqlx包用函数查询的时候采取的是反射，所以需要大写，让sqlx可以访问到；database那个包小写是因为指定赋值给指定的字段
	ID   int    `db:"id"` //这里用tag是因为查询的db字段名称和struct这里的字段名字不一样
	Name string `db:"name"`
	Age  int    `db:"age"`
}

//queryRowDemo 查询单条数据
func queryRowDemo() {
	sqlStr := "select id , name , age from user where id = ?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed , err:%v\n", err)
		return
	}
	fmt.Printf("id:%d , name:%s , age:%d\n", u.ID, u.Name, u.Age)
}

//queryMultiDemo 查询多条数据
func queryMultiDemo() {
	sqlStr := "select id , name , age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query select failed , err:%v\n", err)
		return
	}
	fmt.Printf("users: %#v\n", users)

}
func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init db failed , err:%v\n", err)
		return
	}
	fmt.Printf("init db success")
	//queryRowDemo()
	queryMultiDemo()
}
