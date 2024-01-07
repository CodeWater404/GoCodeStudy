package logic

import (
	"context"
	"go-zero/feng/rpc_study/user_api_rpc/rpc/types/user"

	"go-zero/feng/rpc_study/user_api_rpc/api/internal/svc"
	"go-zero/feng/rpc_study/user_api_rpc/api/internal/types"

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

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	response, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId: uint32(req.ID),
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		UserId:    int(response.UserId),
		UserrName: response.Username,
	}, nil
}
