package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

/*
*

	  @author: CodeWater
	  @since: 2023/11/6
	  @desc: crud
		https://www.liwenzhou.com/posts/Go/mysql/

*
*/
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

// queryRowDemo 查询单条数据
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

// queryMultiRowDemo 查询多条数据
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

// insertRowDemo 差入单条数据
func insertRowDemo() {
	sqlStr := "insert into user(name , age) values(? , ?)"
	ret, err := db.Exec(sqlStr, "water", 22)
	if err != nil {
		fmt.Printf("insert failed , err:%v\n", err)
		return
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert id failed , err:%v\n", err)
		return
	}
	fmt.Printf("insert success , the id is %d\n", theId)
}

// updateRowDemo 更新一条数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 30, 3)
	if err != nil {
		fmt.Printf("update failed , err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed , err:%v\n", err)
		return
	}
	fmt.Printf("udpate success , affected rows:%v\n", n)
}

// deleteRowDemo 删除单条数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed , err:%v\n", err)
		return
	}
	fmt.Printf("delete success , affected rows:%v\n", n)
}

/*
prepareQueryDemo 预处理查询
先发送sql语句给mysql，然后再发送数据；之前的操作都是先把占位符替换掉然后发送sql直接执行。
场景：
优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
避免SQL注入问题。
*/
func prepareQueryDemo() {
	sqlStr := "select id , name , age from user where id > ?"
	ret, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("db prepare failed , err:%v\n", err)
		return
	}
	defer ret.Close()
	rows, err := ret.Query(0)
	if err != nil {
		fmt.Printf("prepare query failed , err:%v\n", err)
		return
	}
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("prepare scan failed , err:%v\n", err)
			return
		}
		fmt.Printf("id:%d , name:%s , age:%d\n", u.id, u.name, u.age)
	}
}

// prepareInsertDemo 预处理差入多条数据
func prepareInsertDemo() {
	sqlStr := "insert into user(name , age) values(? , ?)"
	ret, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("db prepare failed , err:%v\n", err)
		return
	}
	defer ret.Close()
	_, err = ret.Exec("water", 22)
	if err != nil {
		fmt.Printf("prepare 1 insert failed , err:%v\n", err)
		return
	}
	_, err = ret.Exec("lion", 25)
	if err != nil {
		fmt.Printf("prepare 2 insert failed , err:%v\n", err)
		return
	}
	fmt.Println("prepare insert success")
}

// transactionDemo 模拟事务
func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed , err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id = ?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql1 failed , err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret1.RowAffected() failed , err:%v\n", err)
		return
	}
	sqlStr2 := "Update user set age=40 where id = ?"
	ret2, err := tx.Exec(sqlStr2, 4)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql2 failed , err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret2.RowsAffected failed , err:%v\n", err)
		return
	}
	fmt.Println(affRow1, affRow2)

	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("commit transaction!")
		tx.Commit()
	} else {
		tx.Rollback()
		fmt.Println("rollback transaction ...")
	}
	fmt.Println("exec trans success!!")

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
	//queryMultiRowDemo()
	//insertRowDemo()
	//updateRowDemo()
	//deleteRowDemo()
	//prepareQueryDemo()
	//prepareInsertDemo()
	transactionDemo()
}
