package main

import (
	"learngo/errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/4/16
  @desc: 服务器统一错误处理示例:
		defer + panic +recover
	PS: add library: _ "net/http/pprof"
		1. can access this address:http://localhost:8888/debug/pprof/
		2. can use this command: go tool pprof http://localhost:8888/debug/pprof/profile
		to look at the cpu information for 30 seconds
	(see library file comments for more information)

**/

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		//错误panic处理
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		//处理接口返回的报错类型
		if err != nil {
			//加log打印错误请求日志
			log.Printf("error handing request:%s", err.Error())
			//fatalf会结束程序的执行。
			//log.Fatalf("error handling request: %s", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)

				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	//浏览器访问地址：http://localhost:8888/list/fib.txt
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
