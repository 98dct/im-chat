package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"im-chat/apps/user/models"
	"im-chat/apps/user/rpc/internal/config"
	"im-chat/pkg/constants"
	"im-chat/pkg/ctxdata"
	"time"
)

type ServiceContext struct {
	Config config.Config

	*redis.Redis
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:     c,
		Redis:      redis.MustNewRedis(c.Redisx),
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}

func (svc *ServiceContext) SetRootToken() error {

	// 生成jwt
	systemToken, err := ctxdata.GetJwtToken(svc.Config.Jwt.AccessSecret, time.Now().Unix(),
		99999999, constants.SYSTEM_ROOT_UID)
	if err != nil {
		return err
	}
	// 写入redis
	return svc.Redis.Set(constants.REDIS_SYSTEM_ROOT_TOKEN, systemToken)
}
