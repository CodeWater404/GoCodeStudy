package logic

import (
	"context"
	"go-zero/feng/model_study/user_gorm/models"

	"go-zero/feng/model_study/user_gorm/internal/svc"
	"go-zero/feng/model_study/user_gorm/internal/types"

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
	//create table
	//err = l.svcCtx.DB.Create(&models.UserModel{
	//	Username: "code2",
	//	Password: "222",
	//}).Error
	//fmt.Println("====>create table:", err)
	//return "", err

	var user models.UserModel
	err = l.svcCtx.DB.Take(&user, "username=? and password=?", req.Username, req.Password).Error
	if err != nil {
		return "", err
	}

	return user.Username, nil
}
