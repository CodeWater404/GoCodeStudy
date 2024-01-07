package usercreatelogic

import (
	"context"
	"fmt"

	"go-zero/feng/rpc_study/user_group/rpc/internal/svc"
	"go-zero/feng/rpc_study/user_group/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	fmt.Println("===>rpc group:", in.Username, in.Password)

	return &user.UserCreateResponse{}, nil
}
