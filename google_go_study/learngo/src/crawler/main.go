package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/5/21
  @desc: $
**/

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}
