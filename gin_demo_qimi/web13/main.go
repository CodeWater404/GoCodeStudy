package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/6/2
  @desc: 获取请求的path（url）上的参数
	需要注意url的冲突:
	1. /:name/:age
	2. /blog/:year/:month   (这个路由其实是包括上一个路由的，但是现在的版本1.9.1好像不会报错了)
**/

func main() {
	r := gin.Default()

	r.GET("/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("=====>>>> gin run failed , err:%v\n", err)
		return
	}
}
