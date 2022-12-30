package svc

import (
	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	rds := c.RedisConf.NewRedis()
	if !rds.Ping() {
		logx.Error("initialize redis failed")
		return nil
	}

	c.GlobalEnv.NewGlobalEnv()

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rds,
	}
}
