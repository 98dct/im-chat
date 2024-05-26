package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"im-chat/apps/user/models"

	"im-chat/apps/user/rpc/internal/svc"
	"im-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	// todo: add your logic here and delete this line

	//根据用户名、电话号码和id独立查询
	var (
		userEntitys = make([]*models.Users, 0)
		err         error
	)

	if in.Phone != "" {
		userEntity, err := l.svcCtx.FindOneByPhone(l.ctx, in.Phone)
		if err == nil {
			userEntitys = append(userEntitys, userEntity)
		}
	} else if in.Name != "" {
		userEntitys, err = l.svcCtx.ListByName(l.ctx, in.Name)
	} else if len(in.Ids) > 0 {
		userEntitys, err = l.svcCtx.ListByIds(l.ctx, in.Ids)
	}

	if err != nil {
		fmt.Println("1" + err.Error())
		return nil, err
	}

	resp := make([]*user.UserEntity, 0)
	err = copier.Copy(&resp, userEntitys)

	if err != nil {
		fmt.Println("2" + err.Error())
		return nil, err
	}

	return &user.FindUserResp{
		User: resp,
	}, nil
}
