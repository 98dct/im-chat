package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"im-chat/apps/social/models"
	"im-chat/apps/social/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	models.FriendsModel
	models.FriendRequestsModel
	models.GroupsModel
	models.GroupRequestsModel
	models.GroupMembersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config:              c,
		FriendsModel:        models.NewFriendsModel(sqlConn, c.Cache),
		FriendRequestsModel: models.NewFriendRequestsModel(sqlConn, c.Cache),
		GroupsModel:         models.NewGroupsModel(sqlConn, c.Cache),
		GroupRequestsModel:  models.NewGroupRequestsModel(sqlConn, c.Cache),
		GroupMembersModel:   models.NewGroupMembersModel(sqlConn, c.Cache),
	}
}
