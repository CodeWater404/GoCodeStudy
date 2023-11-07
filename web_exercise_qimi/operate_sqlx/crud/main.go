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

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init db failed , err:%v\n", err)
		return
	}
	fmt.Printf("init db success")
}
