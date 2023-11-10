package main

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

//func main() {
//	InitLogger()
//	defer logger.Sync()
//	simpleHttpGet("www.baidu.com")
//	simpleHttpGet("http://www.baidu.com")
//}

func main() {
	InitLogger()
	//使用自定义gin中间件
	r := gin.New()
	r.Use(GinLogger(logger), GinRecovery(logger, true))
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	r.Run()
}
