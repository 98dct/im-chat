package svc

import (
	"im-chat/apps/im/immodels"
	"im-chat/apps/im/ws/internal/config"
)

type ServiceContext struct {
	Config config.Config
	immodels.ChatLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ChatLogModel: immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}
}
