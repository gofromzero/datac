package svc

import (
	"datac/internal/config"
	"datac/third_party/gormcli"
	"datac/third_party/miniocli"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	MysqlDB  *gorm.DB
	MinioCli *miniocli.MinioCli
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		MysqlDB:  gormcli.NewGorm(c.MysqlDSN),
		MinioCli: miniocli.NewMinio(c.Minio.Endpoint, c.Minio.AccessKeyID, c.Minio.SecretAccessKey),
	}
}
