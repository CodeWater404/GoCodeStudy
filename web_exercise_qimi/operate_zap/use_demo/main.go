package main

import (
	"go.uber.org/zap"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/11/9
  @desc: zap的基本使用
**/

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("error fetching url...", zap.String("url", url), zap.Error(err))
	} else {
		logger.Info("success!!!", zap.String("statusCode", resp.Status), zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}
