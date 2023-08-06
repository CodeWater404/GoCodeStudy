package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/8/6
  @desc: 用context实现一个客户端超时控制的案例
**/

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 5)
		_, _ = fmt.Fprintf(w, "slow response....")
		return
	}
	_, _ = fmt.Fprintf(w, "quick response!!!!!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
