package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/11/6
  @desc: crud
**/
var (
	db  *sql.DB
	dsn = "root:%s@tcp(%s:3306)/gin_qimi?charset=utf8mb4&parseTime=True"
)

func initDB() (err error) {
	ten, tenDb := os.Getenv("tencent"), os.Getenv("tencent_mysql")
	if ten == "" || tenDb == "" {
		return errors.New("get env failed, ten:" + ten + ", tenDb:" + tenDb)
	}
	dsn = fmt.Sprintf(dsn, ten, tenDb)

	//这里不会校验账号和密码
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	//尝试和db建立连接，校验账号和密码
	err = db.Ping()
	if err != nil {
		return nil
	}

	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("db connect fail: %v\n", err)
	}
	defer db.Close()
	fmt.Println("db connecnt success!")

}
