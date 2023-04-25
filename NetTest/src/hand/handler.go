package hand

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/4/24
  @desc: 统一的业务逻辑处理,代理请求到服务器b
**/

/**ConnectServer
** @Description: http
** @param w
** @param r
** @return error
**/
func ConnectServer(w http.ResponseWriter, r *http.Request) error {
	conn, err := net.Dial("tcp", "127.0.0.1:9222")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return errors.New("connect the remote server failed......")
	}
	defer conn.Close()

	// Send the HTTP request to the remote server
	r.Write(conn)

	// Read the response from the remote server
	resp, err := http.ReadResponse(bufio.NewReader(conn), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return errors.New("the remote server reply failed......")
	}
	defer resp.Body.Close()

	// Read the response body
	macAddr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return errors.New("io read failed......")
	}

	// 修改响应体
	newBody := strings.Replace(string(macAddr), string(macAddr), fmt.Sprintf("http B mac地址是：%s", macAddr), -1)
	// Set the response body and headers
	_, err = w.Write([]byte(newBody))
	if err != nil {
		return errors.New("write response body failed......")
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Proxy-Server", "GoProxy")

	// Log the request and response
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Printf("=========Http Request time：%v , method: %s , path: %s\nResponse status:%s\n", formattedTime, r.Method, r.URL.Path, resp.Status)
	fmt.Println()

	return nil
}

/**ConnectServerForHttps
** @Description: https
** @param w
** @param r
** @return error
**/
func ConnectServerForHttps(w http.ResponseWriter, r *http.Request) error {
	// 创建一个自定义的 Transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 创建一个使用自定义 Transport 的 Client
	client := &http.Client{Transport: tr}

	// 发起 GET 请求
	resp, err := client.Get("http://127.0.0.1:9222")
	if err != nil {
		fmt.Println("Error: ", err)
		return fmt.Errorf("test(): %s", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	macAddr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return fmt.Errorf("test(): %s", err)
	}

	// 修改响应体
	newBody := strings.Replace(string(macAddr), string(macAddr), fmt.Sprintf("https B mac地址是：%s", macAddr), -1)
	_, err = w.Write([]byte(newBody))
	if err != nil {
		return errors.New("io read failed......")
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Proxy-Server", "GoProxy")

	// Log the request and response
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Printf("==========Https Request time：%v , method: %s , path: %s\nResponse status:%s\n", formattedTime, r.Method, r.URL.Path, resp.Status)
	fmt.Println()

	return nil
}
