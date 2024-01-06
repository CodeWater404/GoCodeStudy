package logic

import (
	"context"

	"go-zero/feng/zero_study/user/rpc/internal/svc"
	"go-zero/feng/zero_study/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	return &user.UserResponse{
		Id:     "1",
		Name:   "code",
		Gender: true,
	}, nil
}
