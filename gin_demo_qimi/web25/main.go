package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/6/6
  @desc: 用gin、gorm做一个待办事项的简单项目bubble
**/

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/**initMysql
** @Description: mysql初始化
** @return err
**/
func initMysql() (err error) {
	arg := "root:mysql's password@(ip address:port)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", arg)
	if err != nil {
		fmt.Printf("======>>> gorm connect failed , err:%v\n", err)
		return
	}
	return DB.DB().Ping()
}

func main() {
	//数据库连接
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//加载静态资源
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("/v1")
	{
		//add
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			//从请求中拿出数据
			c.BindJSON(&todo)
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//select all data
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//select one data(这里其实和下面修改一条记录没啥区别，只不过少了保存到数据库那一步操作，而且更重要的一点是，在页面上不会用到这个接口的功能。。。。)
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})
		//modify one data
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//delete one data
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}

	err = r.Run(":9000")
	if err != nil {
		fmt.Printf("=====>>> gin run failed , err:%v\n", err)
		return
	}

}
