// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero/feng/model_study/user_gorm/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/users"),
	)
}