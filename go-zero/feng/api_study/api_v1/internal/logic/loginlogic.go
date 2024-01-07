package logic

import (
	"context"

	"go-zero/feng/api_study/api_v1/internal/svc"
	"go-zero/feng/api_study/api_v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	return &types.LoginResponse{
		Code: 0,
		Data: "api version 1",
		Msg:  "Got messages from login api !",
	}, nil
}
