package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web25/dao"
	"web25/models"
)

/**
  @author: CodeWater
  @since: 2023/6/6
  @desc: 控制层，获取请求的操作
**/
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	//从请求中拿出数据
	c.BindJSON(&todo)
	if err := dao.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	var todoList []models.Todo
	if err := dao.DB.Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func GetATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "get a todo item failed! Please retry...."})
	}
	var todo models.Todo
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	var todo models.Todo
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	c.BindJSON(&todo)
	if err := dao.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
	if err := dao.DB.Where("id=?", id).Delete(models.Todo{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
