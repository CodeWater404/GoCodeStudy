package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero/feng/api_study/api_jwt/internal/svc"
	"go-zero/feng/api_study/api_jwt/internal/types"
	"go-zero/feng/common/jwt"
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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	auth := l.svcCtx.Config.Auth
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		UserName: "code",
		Role:     1,
		Testhhh:  "testhhh from login logic",
	}, auth.AccessSecret, auth.AccessExpire)

	if err != nil {
		return "", err
	}

	return token, err
}
