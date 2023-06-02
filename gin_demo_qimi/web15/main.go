package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

/**
  @author: CodeWater
  @since: 2023/6/2
  @desc: gin: 文件上传（单个文件示例，多个for循环）
**/

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		//从前端那里拿的，定义的文件的变量名
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			dst := path.Join("./", f.Filename)
			_ = c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Printf("=====>>> gin run failed , err: %v\n", err)
		return
	}
}
