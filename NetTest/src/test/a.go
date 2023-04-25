package main

import (
	"NetTest/hand"
	"log"
	"net/http"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/4/23
  @desc:
	todo：
		https报错：
			2023/04/26 02:08:26 http: TLS handshake error from 127.0.0.1:58055: remote error: tls: unknown certificate
			2023/04/26 02:08:26 http: TLS handshake error from 127.0.0.1:58056: remote error: tls: unknown certificate

**/

type userError interface {
	error
	Message() string
}

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

func main() {
	// 第一个 HTTP 服务器
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", errWrapper(hand.ConnectServer))
	server1 := &http.Server{
		Addr:    ":8855",
		Handler: mux1,
	}

	// 第二个 HTTP 服务器
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", errWrapper(hand.ConnectServerForHttps))
	server2 := &http.Server{
		Addr:    ":8854",
		Handler: mux2,
	}

	// 启动两个服务器
	go server1.ListenAndServe()
	go server2.ListenAndServeTLS("server.crt", "server.key")

	// 防止程序退出
	select {}
}
