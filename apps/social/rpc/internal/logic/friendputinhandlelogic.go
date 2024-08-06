package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"im-chat/apps/social/models"
	"im-chat/pkg/constants"
	"im-chat/pkg/xerr"

	"im-chat/apps/social/rpc/internal/svc"
	"im-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrFriendReqBeforePass   = xerr.NewMsg("好友申请已经通过！")
	ErrFriendReqBeforeRefuse = xerr.NewMsg("好友申请已经被拒绝！")
)

type FriendPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {
	// todo: add your logic here and delete this line

	// 获得好友申请记录
	friendRequests, err := l.svcCtx.FriendRequestsModel.FindOne(l.ctx, uint64(in.FriendReqId))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "get frinedRequest serr: %v, req:%v",
			err, in.FriendReqId)
	}
	// 验证是否有处理，不能重复处理
	switch constants.HandlerResult(friendRequests.HandleResult.Int64) {
	case constants.PassHandlerResult:
		return nil, errors.WithStack(ErrFriendReqBeforePass)
	case constants.RefuseHandlerResult:
		return nil, errors.WithStack(ErrFriendReqBeforeRefuse)
	}

	//修改申请结果并建立两条还有关系记录 用事务！！ 同步数据
	friendRequests.HandleResult.Int64 = int64(in.HandleResult)

	err = l.svcCtx.FriendRequestsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		if err := l.svcCtx.FriendRequestsModel.Update(ctx, session, friendRequests); err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "update friendship fail err: %v, req:%v",
				err, friendRequests)
		}

		if constants.HandlerResult(in.HandleResult) != constants.PassHandlerResult {
			return nil
		}

		friends := []*models.Friends{
			{
				UserId:    friendRequests.UserId,
				FriendUid: friendRequests.ReqUid,
			},
			{
				UserId:    friendRequests.ReqUid,
				FriendUid: friendRequests.UserId,
			},
		}

		_, err := l.svcCtx.FriendsModel.Inserts(ctx, session, friends...)
		if err != nil {
			return errors.Wrapf(xerr.NewDBErr(), "friends insert fail err: %v, req:%v",
				err, friendRequests)
		}

		return nil

	})

	return &social.FriendPutInHandleResp{}, err
}
