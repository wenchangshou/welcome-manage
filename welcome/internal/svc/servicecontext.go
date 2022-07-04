package svc

import (
	"welcome-manage/welcome/internal/config"
	"welcome-manage/welcome/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	TitleModel model.TitleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		TitleModel: model.NewTitleModel(conn, c.CacheRedis),
	}
}
