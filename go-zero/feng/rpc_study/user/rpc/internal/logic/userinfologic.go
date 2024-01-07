package logic

import (
	"context"

	"go-zero/feng/rpc_study/user/rpc/internal/svc"
	"go-zero/feng/rpc_study/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	return &user.UserInfoResponse{
		UserId:   in.UserId,
		Username: "coderpc",
	}, nil
}
