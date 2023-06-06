package routers

import (
	"github.com/gin-gonic/gin"
	"web25/controller"
)

/**
  @author: CodeWater
  @since: 2023/6/6
  @desc: 路由相关的
**/

func SetupRouters() *gin.Engine {
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//加载静态资源
	r.Static("/static", "static")

	r.GET("/", controller.Index)

	v1Group := r.Group("/v1")
	{
		//add
		v1Group.POST("/todo", controller.CreateATodo)
		//select all data
		v1Group.GET("/todo", controller.GetTodoList)
		//select one data(这里其实和下面修改一条记录没啥区别，只不过少了保存到数据库那一步操作，而且更重要的一点是，在页面上不会用到这个接口的功能。。。。)
		v1Group.GET("/todo/:id", controller.GetATodo)
		//modify one data
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//delete one data
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
