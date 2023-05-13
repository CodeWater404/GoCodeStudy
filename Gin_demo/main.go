package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/5/11
  @desc: $
**/

/**myHandler
** @Description:自定义go中间件（类似与java中的拦截器，实现在请求之前）
**/
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		//通过自定义的中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
		context.Set("usersession", "userid-1")
		context.Next() //放行
		//context.Abort() //阻断
	}
}

func main() {
	//创建一个服务
	ginServer := gin.Default()
	//设置网页标签上的图标
	ginServer.Use(favicon.New("./favicon.ico"))

	//4.加载静态页面（全局加载）
	ginServer.LoadHTMLGlob("templates/*")
	/*//加载指定文件
	ginServer.LoadHTMLFiles("templates/index.html")*/

	//5.加载资源文件
	ginServer.Static("/static", "./static")

	//1.访问地址，处理我们的请求 Request Response
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})

	/*//2.restful api
	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "post user"})
	})
	ginServer.PUT("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "put user"})
	})
	ginServer.DELETE("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "delete user"})
	})*/

	//3.响应一个页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是go后台传递过来的数据",
		})
	})

	//6. 接收前端传递过来的参数
	//6.1 url?userid=xxx&username=codewater
	ginServer.GET("/user/info", myHandler(), func(context *gin.Context) {
		//调用，取出中间件中的值
		userSession := context.MustGet("usersession").(string) //转成string
		log.Println("=====================>", userSession)

		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"useridQuery": userid,
			"username":    username,
		})
	})

	//6.2 url/user/info/1/codewater
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"useridParam": userid,
			"username":    username,
		})
	})

	//7.前端给后端传json（序列化）
	ginServer.POST("/json", func(context *gin.Context) {
		// request body
		data, _ := context.GetRawData()

		var m map[string]interface{}
		//包装为json数据
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})

	//8. 处理表单数据
	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")

		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	//9. 路由
	ginServer.GET("/test", func(context *gin.Context) {
		//重定向:只有状态码正确才会转向
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	//404 no route(输入不存在的路由时就会返回404页面)
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	//路由组:用户组
	userGroup := ginServer.Group("/user")
	{
		//这个时候的请求url路径就是：url/user/add
		userGroup.GET("/add")
		userGroup.POST("/login")
		userGroup.POST("/logout")
	}
	//订单组
	orderGroup := ginServer.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.DELETE("/delete")
	}

	//运行，服务器端口
	ginServer.Run(":8082")
}
