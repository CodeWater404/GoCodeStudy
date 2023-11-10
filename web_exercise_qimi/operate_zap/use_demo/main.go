package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/11/9
  @desc: zap的基本使用
	这个日志程序中唯一缺少的就是日志切割归档功能。
	https://www.liwenzhou.com/posts/Go/zap/#c-1-3-3
**/

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//logger, _ = zap.NewProduction() //default config

	//custom logger's setting
	/*todo:未理解这个addCallerSkip
	当我们不是直接使用初始化好的logger实例记录日志，而是将其包装成一个函数等，此时日录日志的函数调用链会增加，
	想要获得准确的调用信息就需要通过AddCallerSkip函数来跳过。
	*/
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	sugarLogger = logger.Sugar()
}

//func getLogWriter() zapcore.WriteSyncer {
//	file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
//	return file
//}

/*
getLogWriter
使用lumberjack库进行日志切割归档
*/
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log", //Filename: 日志文件的位置
		MaxSize:    1,            //MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 2,            //MaxBackups：保留旧文件的最大个数
		MaxAge:     30,           //MaxAges：保留旧文件的最大天数
		Compress:   false,        //Compress：是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	//return zapcore.NewJSONEncoder(encoderConfig) //输出json格式
	return zapcore.NewConsoleEncoder(encoderConfig) //输出终端格式
}

func simpleHttpGet(url string) {
	//sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("error fetching url...", zap.String("url", url), zap.Error(err))
		//sugarLogger.Errorf("error fetching url %s : error = %s", url, err)
	} else {
		logger.Info("success!!!", zap.String("statusCode", resp.Status), zap.String("url", url))
		//sugarLogger.Infof("success!!! statusCode = %s for url %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}
