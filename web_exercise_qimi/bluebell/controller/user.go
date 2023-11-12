package controller

import (
	"fmt"
	"net/http"
	"web_exercise_qimi/bluebell/logic"
	"web_exercise_qimi/bluebell/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

/**
  @author: CodeWater
  @since: 2023/11/12
  @desc: $
**/

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(), //不是validator.ValidationErrors类型就返回原来的错误
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			//一般不暴露内部系统的错误
			"msg": removeTopStruct(errs.Translate(trans)), //是validator.ValidationErrors类型就翻译成中文
		})
		return
	}
	//手动对参数进行校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}

	fmt.Printf("signup===>p: %#v\n", p)
	// 2. 业务处理
	logic.SignUp(p)
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
