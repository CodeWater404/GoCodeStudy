package controller

import (
	"fmt"
	"reflect"
	"strings"
	"web_exercise_qimi/bluebell/models"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

/*
*

	@author: CodeWater
	@since: 2023/11/12
	@desc: $

*
*/

// 定义一个全局的翻译器T
var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	//修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fid reflect.StructField) string {
			//
			name := strings.SplitN(fid.Tag.Get("json"), ",", 2)[0]
			//
			if name == "-" {
				return ""
			}
			return name
		})

		// 为SignUpParam注册自定义校验方法
		v.RegisterStructValidation(SignUpParamStructLevelValidation, models.ParamSignUp{})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		uni := ut.New(enT, zhT, enT)
		// locale通常取决于http请求'Accept-language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator failed , err:%v\n", err)
		}

		switch locale {
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// removeTopStruct 去除掉错误提示中的结构体名称标识
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// SignUpParamStructLevelValidation 自定义SignUpParam的结构体校验
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.ParamSignUp)
	if su.Password != su.RePassword {
		//输出错误提示信息，最后一个参数是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
