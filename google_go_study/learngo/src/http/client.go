package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

/**
  @author: CodeWater
  @since: 2023/5/4
  @desc: http standard library
	how to see official document:
		1.use command to build a server: gotool doc :8888
		2.access address: localhost:8888

**/

func main() {
	//use http to controlthe request header
	//request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	//request.Header.Add("User-Agent", "")
	//resp2, err := http.DefaultClient.Do(request)

	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
