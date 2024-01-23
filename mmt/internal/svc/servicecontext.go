package svc

import (
	"mmt/mmt/internal/config"
	"mmt/mmt/internal/middleware"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redis.Redis
	Mysql   *gorm.DB
	JwtAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Redis:   redis.MustNewRedis(c.Redis),
		Mysql:   NewMysql(c.Mysql.DataSource),
		JwtAuth: middleware.NewJwtAuthMiddleware(c.Auth.AccessSecret).Handle,
	}
}
