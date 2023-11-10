package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
)

/**
  @author: CodeWater
  @since: 2023/11/7
  @desc: crud with sqlx
	https://www.liwenzhou.com/posts/Go/sqlx/
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

// queryRowDemo 查询单条数据
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

// queryMultiDemo 查询多条数据
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

// insertRowDemo 插入一条数据
func insertRowDemo() {
	sqlStr := "insert into user (name , age) values(? , ?)"
	ret, err := db.Exec(sqlStr, "cat", 28)
	if err != nil {
		fmt.Printf("insert failed , err:%v\n", err)
		return
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsertId failed , err:%v\n", err)
		return
	}
	fmt.Printf("insert success , the id is %d\n", theId)
}

// updateRowDemo 更新一条数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 66, 6)
	if err != nil {
		fmt.Printf("update failed , err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected failed , err:%v\n", err)
		return
	}
	fmt.Printf("update success , affected rows:%d\n", n)
}

// deleteRowDemo 删除一条数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed , err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected failed , err:%v\n", err)
		return
	}
	fmt.Printf("delete success , affected rows:%d\n", n)
}

// insertUserDemo namedExec绑定sql语句与结构体或map中的同名字段
func insertUserDemo() (err error) {
	sqlStr := "insert into user(name , age) values(:name , :age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "dog",
			"age":  "32",
		})
	return
}

// namedQuery namedQuery绑定sql语句与结构体或map中的同名字段,不过这里是支持查询的
func namedQuery() {
	sqlStr := "select * from user where name=:name"
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{
		"name": "code",
	})
	if err != nil {
		fmt.Printf("namedQuery failed , err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed , err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
	u := &user{
		Name: "qimi",
	}
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("namedQuery failed , err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed , err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

// transactionDemo 事务例子
func transactionDemo() (err error) {
	tx, err := db.Beginx()
	if err != nil {
		fmt.Printf("begin trans failed , err:%v\n", err)
		return err
	}
	//最后对err做处理
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Printf("rollback")
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "update user set age=20 where id = ?"
	rs, err := tx.Exec(sqlStr1, 7)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := "update user set age=50 where id = ?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr2 failed")
	}
	return err

}

// BatchInsertUsers 批量插入，手动拼接sql
func BatchInsertUsers(users []*user) error {
	valueString := make([]string, 0, len(users))
	valueArgs := make([]interface{}, 0, len(users)*2)
	for _, u := range users {
		valueString = append(valueString, "(? , ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	stmt := fmt.Sprintf("insert into user (name , age) values %s", strings.Join(valueString, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

/*Value 使用sqlx库实现批量插入，需要实现下面的value接口
 */
func (u user) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"insert into user (name , age) values(?) , (?) , (?)", users...)
	fmt.Println("batch insert2 sql:" + query)
	fmt.Printf("batch insert2 args:%v\n", args)
	_, err := db.Exec(query, args...)
	return err
}

// BatchInsertUsers3 sqlx内置接口批量插入
func BatchInsertUsers3(users []*user) error {
	_, err := db.NamedExec("insert into user (name , age) values(:name , :age)", users)
	return err
}

// QueryByIDs 批量查询
func QueryByIDs(ids []int) (users []user, err error) {
	query, args, err := sqlx.In("select id ,name , age from user where id in (?)", ids)
	if err != nil {
		return
	}
	//sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}

// QueryAndOrderByIds 批量查询，这里通过FIND_IN_SET保留了按照查询顺序来返回结果；另外一种处理方式是手动自己去排序，
func QueryAndOrderByIds(ids []int) (users []user, err error) {
	strIds := make([]string, 0, len(ids))
	for _, id := range ids {
		strIds = append(strIds, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("select id ,name , age from user where id in (?) order by FIND_IN_SET(id , ?)", ids, strings.Join(strIds, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init db failed , err:%v\n", err)
		return
	}
	fmt.Printf("init db success\n")
	//queryRowDemo()
	//queryMultiDemo()
	//insertRowDemo()
	//updateRowDemo()
	//deleteRowDemo()

	//err := insertUserDemo()
	//if err != nil {
	//	fmt.Printf("insert failed , err:%v\n", err)
	//}else {
	//	fmt.Printf("insert success")
	//}

	//namedQuery()

	//if err := transactionDemo(); err != nil {
	//	fmt.Printf("trans failed , err:%v\n", err)
	//} else {
	//	fmt.Printf("trans success")
	//}

	//u1 := user{Name: "duck", Age: 41}
	//u2 := user{Name: "tiger", Age: 42}
	//u3 := user{Name: "cow", Age: 43}
	//users := []*user{&u1, &u2, &u3}
	//if err := BatchInsertUsers(users); err != nil {
	//	fmt.Printf("BatchInsertUsers failed , err:%v\n", err)
	//}

	//user2 := []interface{}{u1, u2, u3} //元素类型为 interface{}。interface{} 是空接口，
	//if err := BatchInsertUsers2(user2); err != nil {
	//	fmt.Printf("batch insert2 failed , err:%v\n", err)
	//}

	//if err := BatchInsertUsers3(users); err != nil {
	//	fmt.Printf("batch insert3 failed , err:%v\n", err)
	//}

	//ids := []int{1, 2, 5, 6}
	//users2, err := QueryByIDs(ids)
	//if err != nil {
	//	fmt.Printf("queryByIds failed , err:%v\n", err)
	//}
	//fmt.Printf("queryByIds:%v\n", users2)

	ids := []int{10, 8, 1, 2, 5}
	users2, err := QueryAndOrderByIds(ids)
	if err != nil {
		fmt.Printf("QueryAndOrderByIds failed , err:%v\n", users2)
	}
	fmt.Printf("QueryAndOrderByIds:%v\n", users2)

}
