// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-zero/feng/rpc_study/user_gorm/rpc/internal/logic"
	"go-zero/feng/rpc_study/user_gorm/rpc/internal/svc"
	"go-zero/feng/rpc_study/user_gorm/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServer) UserCreate(ctx context.Context, in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	l := logic.NewUserCreateLogic(ctx, s.svcCtx)
	return l.UserCreate(in)
}
