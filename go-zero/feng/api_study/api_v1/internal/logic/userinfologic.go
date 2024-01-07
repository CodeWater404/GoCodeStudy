package logic

import (
	"context"

	"go-zero/feng/api_study/api_v1/internal/svc"
	"go-zero/feng/api_study/api_v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	return &types.UserInfoResponse{
		Code: 0,
		Data: "api version 1",
		Msg:  "Got message from userInfo api!!",
	}, nil
}
