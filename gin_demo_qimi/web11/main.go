package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/6/1
  @desc: 获取前端链接上面的参数:query-string
	1. 几种写法
	2，多个值获取
**/

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//前端传递的一切东西都可以在c中找到(如果有多个值，可以查多次；再多的话需要用其他方法查询)
		name := c.Query("query")
		age := c.Query("age")
		//如果查询不到，那就使用默认值
		other := c.DefaultQuery("other", "hhhhhhh")
		//查询不到，返回值多一个错误的方法
		name2, ok := c.GetQuery("name2")
		if !ok {
			name2 = "query failed , name2......"
		}
		c.JSON(http.StatusOK, gin.H{
			"name":  name,
			"age":   age,
			"other": other,
			"name2": name2,
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("=========>>> gin run falied , err: %v\n ", err)
		return
	}
}
