package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

/**
  @author: CodeWater
  @since: 2023/4/23
  @desc:
	该程序将在A服务器上启动两个HTTP服务器，一个监听在8855端口，另一个监听在8854端口。
	当收到HTTP请求时，它将建立一个TCP连接到B服务器的9222端口，并将请求发送到该服务器。然后，它将等待响应，并从响应中读取响应正文。它将修改响应正文以包含所需的MAC地址，并将其发送回A服务器上的浏览器。它还将添加一个X-Proxy-Server头，指示这是一个代理服务器。它将在控制台上记录请求和响应。
	对于HTTPS请求，该程序将使用TLS Dial函数建立一个加密的TCP连接到B服务器的9222端口，并发送请求。然后，它将等待响应，并从响应中读取响应正文。它将修改响应正文以包含所需的MAC地址，并将其发送回A服务器上的浏览器。它还将添加一个X-Proxy-Server头，指示这是一个代理服务器。它将在控制台上记录请求和响应。
	最后，程序还包括一个getMAC函数，它将返回本地机器上第一个网络接口的MAC地址。
	请注意，本示例程序仅供参考，并可能需要根据您的特定需求进行修改。此外，请确保将IP地址和端口更改为您的实际值，并在A服务器上运行该程序以启动代理服务器。

	todo：
		1. 修改由变量指定域名（httpde）
		2. 错误处理
		3. 发布到服务器
		4. 优化，54变成https的
**/

func connectServer(w http.ResponseWriter, r *http.Request) {
	// Create a new TCP connection to the remote server (B)
	//conn, err := net.Dial("tcp", "10.18.13.101:9222")
	conn, err := net.Dial("tcp", "127.0.0.1:9222")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Send the HTTP request to the remote server
	r.Write(conn)

	// Read the response from the remote server
	resp, err := http.ReadResponse(bufio.NewReader(conn), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Modify the response body to add the MAC address
	mac := getMAC()
	newBody := strings.Replace(string(body), "Hello, world!", fmt.Sprintf("http B mac地址是：%s", mac), -1)

	// Set the modified response body and headers
	w.Write([]byte(newBody))
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Proxy-Server", "GoProxy")

	// Log the request and response
	fmt.Printf("Request method: %s , path: %s\nResponse status:%s\n", r.Method, r.URL.Path, resp.Status)
	fmt.Println()
}

func getMacAddrForHttps(w http.ResponseWriter, r *http.Request) {
	// Create a new TCP connection to the remote server (B)
	//conn, err := tls.Dial("tcp", "10.18.13.101:9222", &tls.Config{})
	conn, err := tls.Dial("tcp", "127.0.0.1:9222", &tls.Config{})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Send the HTTPS request to the remote server
	r.Write(conn)

	// Read the HTTPS response from the remote server
	resp, err := http.ReadResponse(bufio.NewReader(conn), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Modify the response body to add the MAC address
	mac := getMAC()
	newBody := strings.Replace(string(body), "Hello, world!", fmt.Sprintf("https B mac地址是：%s", mac), -1)

	// Set the modified response body and headers
	w.Write([]byte(newBody))
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Proxy-Server", "GoProxy")

	// Log the request and response
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	fmt.Println(resp.Status)
}

func main() {
	// Listen on port 8855 for HTTP requests
	http.HandleFunc("/55/", connectServer)

	err := http.ListenAndServe(":8855", nil)
	if err != nil {
		panic(err)
	}

	// Listen on port 8854 for HTTPS requests
	//http.HandleFunc("/54/", getMacAddrForHttps)

	// Listen on port 8855 and 8854 for HTTP and HTTPS requests, respectively
	//err2 := http.ListenAndServeTLS(":8854", "server.crt", "server.key", nil)
	//if err2 != nil {
	//	panic(err2)
	//}
}

// getMAC returns the MAC address of the first network interface on the local machine
func getMAC() string {
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
