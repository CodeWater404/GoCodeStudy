package logic

import (
	"context"

	"go-zero/code/microhelloworld/mall/user/internal/svc"
	"go-zero/code/microhelloworld/mall/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
