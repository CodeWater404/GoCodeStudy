package logic

import (
	"context"
	"go-zero/feng/model_study/user/api/internal/svc"
	"go-zero/feng/model_study/user/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// insert a data
	//res, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
	//	Id:       1,
	//	Username: "code",
	//	Password: "111",
	//})
	//if err != nil {
	//	return "", err
	//}
	//fmt.Println(res)
	//return "xxx.xxx", nil

	user, err := l.svcCtx.UserModel.FindOneByUsernameAndPassword(l.ctx, req.Username, req.Password)

	if err != nil {
		return "", err
	}
	return user.Username, err
}
