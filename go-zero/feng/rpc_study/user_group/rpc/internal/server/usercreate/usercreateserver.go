// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-zero/feng/rpc_study/user_group/rpc/internal/logic/usercreate"
	"go-zero/feng/rpc_study/user_group/rpc/internal/svc"
	"go-zero/feng/rpc_study/user_group/rpc/types/user"
)

type UserCreateServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserCreateServer
}

func NewUserCreateServer(svcCtx *svc.ServiceContext) *UserCreateServer {
	return &UserCreateServer{
		svcCtx: svcCtx,
	}
}

func (s *UserCreateServer) UserCreate(ctx context.Context, in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	l := usercreatelogic.NewUserCreateLogic(ctx, s.svcCtx)
	return l.UserCreate(in)
}
