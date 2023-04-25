package main

import (
	"fmt"
	"net"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2023/4/23
  @desc: b端web服务
**/

// getMAC returns the MAC address of the first network interface on the local machine
func GetMAC() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				hwaddr := i.HardwareAddr
				return hwaddr.String()
			}
		}
	}
	return ""
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		macAddr := GetMAC()
		if macAddr == "" {
			fmt.Fprintf(w, "http B mac地址获取失败！！！")
			return
		}
		fmt.Fprintf(w, "%s\n", GetMAC())
		return
	}

	fmt.Fprintf(w, "Error: invalid request method")
	return
}

func main() {
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":9222", nil)
	if err != nil {
		panic(err)
	}
}
