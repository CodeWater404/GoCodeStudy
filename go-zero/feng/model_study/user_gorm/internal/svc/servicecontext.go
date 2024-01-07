package svc

import (
	"go-zero/feng/common/init_gorm"
	"go-zero/feng/model_study/user_gorm/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     init_gorm.InitGorm(c.Mysql.DataSource),
	}
}
