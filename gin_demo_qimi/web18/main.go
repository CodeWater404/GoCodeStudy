package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/6/4
  @desc: 中间件
	1. 可以单个请求的时候注册，也可以全局注册
	2， 可以注册多个中间件
**/

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	//获取从中间件中设置的值
	name, ok := c.Get("name")
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"name": name,
	})
}

/**m1
** @Description: 中间件，参数只要是context即可
** @param c
**/
func m1(c *gin.Context) {
	fmt.Println("========>>>>m1 in")
	start := time.Now()

	//注意如果使用协程的话，需要使用c的拷贝，不然后面并发会有数据不一致的情况
	//go funcxx(c.Copy())

	//name, _ := c.Get("name") //name= nil，是在m2中设置的，所以起码等m2执行完成后才能获取到。
	//fmt.Printf("=====>>>>before   m1 get name:%v\n", name)

	//调用后的处理函数
	c.Next()
	//阻止调用后续的处理函数
	//c.Abort()

	name, _ := c.Get("name")
	fmt.Printf("=====>>>>m1 get name:%v\n", name)

	cost := time.Since(start)
	fmt.Printf("========>>>>m1 out , cost: %v\n", cost)
}

func m2(c *gin.Context) {
	fmt.Println("m2 in")
	//跨中间件获取值
	c.Set("name", "codewater")
	c.Next()
	fmt.Println("m2 out")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者一些其它准备工作
	return func(c *gin.Context) {
		fmt.Println("authMiddleware in..........")
		if doCheck {
			//存放具体的逻辑3
			//是否登录的判断
			//if是登录用户
			//c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	//全局注册中间件
	/**2023-6-3
	原来视频中重复注册会失败，但是现在重复注册是可以成功的，并且执行了两遍.
	据我观察报错：应该是up主第二个路由组写错了
	*/
	r.Use(m1, m2, authMiddleware(true))

	//单个请求的中间件注册
	//r.GET("/index", m1, indexHandler)
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	//路由组的中间件注册
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "xxGroup",
			})
		})
	}
	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "xx2Group",
			})
		})
	}

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("=====>>>> gin run failed , err: %v\n", err)
		return
	}
}
