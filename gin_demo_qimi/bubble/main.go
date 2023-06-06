package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"web25/dao"
	"web25/models"
	"web25/routers"
)

/**
  @author: CodeWater
  @since: 2023/6/6
  @desc: 重构项目bubble
**/

func main() {
	//数据库连接
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	//数据绑定
	dao.DB.AutoMigrate(&models.Todo{})

	//注册路由
	r := routers.SetupRouters()
	err = r.Run(":9000")
	if err != nil {
		fmt.Printf("=====>>> gin run failed , err:%v\n", err)
		return
	}

}
