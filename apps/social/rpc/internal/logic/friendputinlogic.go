package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"im-chat/apps/social/models"
	"im-chat/pkg/constants"
	"im-chat/pkg/xerr"
	"time"

	"im-chat/apps/social/rpc/internal/svc"
	"im-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// todo: add your logic here and delete this line

	// 申请人与目标是否已经是好友关系
	friends, err := l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, in.UserId, in.ReqUid)
	if err != nil && err != models.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find frineds by uid and rid err: %v, req:%v",
			err, in)
	}

	if friends != nil {
		return &social.FriendPutInResp{}, err
	}

	// 是否有过申请中、申请是否成功、没有完成
	friendRequests, err := l.svcCtx.FriendRequestsModel.FindByReqUidAndUserId(l.ctx, in.ReqUid, in.UserId)
	if err != nil && err != models.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find frinedRequests by rid and uid err: %v, req:%v",
			err, in)
	}

	if friendRequests != nil {
		return &social.FriendPutInResp{}, err
	}
	// 创建申请记录

	_, err = l.svcCtx.FriendRequestsModel.Insert(l.ctx, &models.FriendRequests{
		UserId: in.UserId,
		ReqUid: in.ReqUid,
		ReqMsg: sql.NullString{
			Valid:  true,
			String: in.ReqMsg,
		},
		ReqTime: time.Unix(in.ReqTime, 0),
		HandleResult: sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert frinedRequests err: %v, req:%v",
			err, in)
	}

	return &social.FriendPutInResp{}, nil
}
