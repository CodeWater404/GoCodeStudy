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
var sugarLogger *zap.SugaredLogger

func InitLogger() {
	logger, _ = zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		//logger.Error("error fetching url...", zap.String("url", url), zap.Error(err))
		sugarLogger.Errorf("error fetching url %s : error = %s", url, err)
	} else {
		//logger.Info("success!!!", zap.String("statusCode", resp.Status), zap.String("url", url))
		sugarLogger.Infof("success!!! statusCode = %s for url %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}
