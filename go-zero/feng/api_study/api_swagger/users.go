package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero/feng/common/response"
	"net/http"

	"go-zero/feng/api_study/api_jwt/internal/config"
	"go-zero/feng/api_study/api_jwt/internal/handler"
	"go-zero/feng/api_study/api_jwt/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(JwtUnauthorizedResult))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// JwtUnauthorizedResult 鉴权失败后的回调处理，详细逻辑还需要再细化，这只是示例
func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println(err)
	httpx.WriteJson(w, http.StatusOK, response.Body{10086, "", "鉴权失败"})
}
