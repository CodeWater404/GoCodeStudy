package logic

import (
	"context"

	"go-zero/feng/api_study/api_v2/internal/svc"
	"go-zero/feng/api_study/api_v2/internal/types"

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
		UserId:    0,
		UserrName: "code water",
	}, nil
}
