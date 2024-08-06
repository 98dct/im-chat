package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"im-chat/apps/social/models"
	"im-chat/pkg/constants"
	"im-chat/pkg/wuid"
	"im-chat/pkg/xerr"

	"im-chat/apps/social/rpc/internal/svc"
	"im-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupCreateLogic) GroupCreate(in *social.GroupCreateReq) (*social.GroupCreateResp, error) {
	// todo: add your logic here and delete this line
	groups := &models.Groups{
		Id:         wuid.GenUid(l.svcCtx.Config.Mysql.Datasource),
		Name:       in.Name,
		Icon:       in.Icon,
		CreatorUid: in.CreatorUid,
		IsVerify:   false,
	}

	err := l.svcCtx.GroupsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		_, err := l.svcCtx.GroupsModel.Insert(l.ctx, session, groups)
		if err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "insert group err %v req %v", err, in)
		}

		// 创建组的人，自动加入组
		_, err = l.svcCtx.GroupMembersModel.Insert(l.ctx, session, &models.GroupMembers{
			GroupId:   groups.Id,
			UserId:    in.CreatorUid,
			RoleLevel: int64(constants.CreatorGroupRoleLevel),
		})
		if err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "insert group member err %v req %v", err, in)
		}
		return nil
	})

	return &social.GroupCreateResp{
		Id: groups.Id,
	}, err
}
