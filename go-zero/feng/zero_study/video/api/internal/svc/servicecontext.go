package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero/feng/zero_study/user/rpc/userclient"
	"go-zero/feng/zero_study/video/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
