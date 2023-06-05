package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/6/5
  @desc: gorm一些默认的设置
**/

type User struct {
	gorm.Model
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time
	Email    string `gorm:"type:varchar(100);unique_index"` //tag
	Role     string `gorm:"size:255"`
	//在数据库中一般都是用下划线分隔开，都是小写
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"` //自增类型
	Address      string  `gorm:"index:addr"`     //会创建一个addr的索引
	IgnoreMe     int     `gorm:"-"`              //忽略本字段
}

type Animal struct {
	//自定义主键，一般不指定就是id字段
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	//自定义列名
	Age int64 `gorm:"column:animal_age"`
}

/**TableName
** @Description: 自定义表名，如果不指定那就是表名加s.这里只要实现这个函数，那么就会自动执行，即使main中没有调用
** @receiver Animal
** @return string
**/
func (Animal) TableName() string {
	return "testAnimal"
}

func main() {
	//修改默认表名规则，这里会给表名加上前缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "prefix_" + defaultTableName
	//}

	args := "root:mysql's password@(ip address:port)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", args)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 表名禁用复数
	db.SingularTable(true)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})
	//另外一种创建表名的方式.
	//db.Table("hhhh_user").CreateTable(&User{})
}
