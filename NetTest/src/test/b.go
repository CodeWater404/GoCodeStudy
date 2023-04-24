package main

import (
	"fmt"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/4/23
  @desc: b端web服务
**/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	err := http.ListenAndServe(":9222", nil)
	if err != nil {
		panic(err)
	}
}
