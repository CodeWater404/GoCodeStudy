package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"go-zero/feng/api_study/api_jwt/internal/svc"
	"go-zero/feng/api_study/api_jwt/internal/types"

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
	userId := l.ctx.Value("user_id").(json.Number)
	fmt.Printf("userId:%v , %T\n", userId, userId)
	username := l.ctx.Value("username").(string)
	uid, _ := userId.Int64()
	//todoi: 属性名和tag都是testhhh，为什么ctx中存行的是testhh？？？？在哪里设置到上下文的？？？？
	tt := l.ctx.Value("testhh").(string)
	fmt.Printf("testhh:%v,%T\n", tt, tt)
	fmt.Printf("===> l ctx:%#+v , svcctx:%v\n", l.ctx, l.svcCtx)

	return &types.UserInfoResponse{
		UserId:    int(uid),
		UserrName: username,
		Testhhh:   tt,
	}, nil
}
