package logic

import (
	"context"
	"errors"
	"go-zero/feng/rpc_study/user_api_rpc/rpc/types/user"

	"go-zero/feng/rpc_study/user_api_rpc/api/internal/svc"
	"go-zero/feng/rpc_study/user_api_rpc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.UserCreateRequest) (resp string, err error) {
	response, err := l.svcCtx.UserRpc.UserCreate(l.ctx, &user.UserCreateRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}

	if response.Err != "" {
		return "", errors.New(response.Err)
	}

	return
}
