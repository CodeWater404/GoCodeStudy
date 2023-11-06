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
	dsn = fmt.Sprintf(dsn, tenDb, ten)

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

	//与数据库连接最大数
	db.SetMaxOpenConns(1)
	//最大空闲连接
	db.SetMaxIdleConns(1)
	return nil
}

type user struct {
	id   int64
	age  int64
	name string
}

//queryRowDemo 查询单条数据
func queryRowDemo() {
	sqlStr := "select id , name , age from user where id = ?"
	var u user
	//非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放(因为queryrow里面没有释放连接，scan里面释放了)
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	/*可以实验：去掉上面的scan
	再加一行
	db.QueryRow(sqlStr , 2)
	然后调用这个函数的时候会发现一直在运行中没有释放连接*/
	if err != nil {
		fmt.Printf("scan failed , err: %v\n", err)
		return
	}
	fmt.Printf("====>func queryRowDemo,id:%v , name:%v , age:%v\n", u.id, u.name, u.age)
}

//queryMultiRowDemo 查询多条数据
func queryMultiRowDemo() {
	sqlStr := "select id , name , age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query multi data failed , err:%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("query multi scan failed , err:%v\n", err)
			return
		}
		fmt.Printf("query multi id:%v , name:%v , age:%v\n", u.id, u.name, u.age)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("db connect fail: %v\n", err)
	}
	defer db.Close()
	fmt.Println("db connect success!")
	fmt.Println("==========================================================")
	//queryRowDemo()
	queryMultiRowDemo()
}
