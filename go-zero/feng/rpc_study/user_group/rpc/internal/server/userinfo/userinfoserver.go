// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-zero/feng/rpc_study/user_group/rpc/internal/logic/userinfo"
	"go-zero/feng/rpc_study/user_group/rpc/internal/svc"
	"go-zero/feng/rpc_study/user_group/rpc/types/user"
)

type UserInfoServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserInfoServer
}

func NewUserInfoServer(svcCtx *svc.ServiceContext) *UserInfoServer {
	return &UserInfoServer{
		svcCtx: svcCtx,
	}
}

func (s *UserInfoServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := userinfologic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}
