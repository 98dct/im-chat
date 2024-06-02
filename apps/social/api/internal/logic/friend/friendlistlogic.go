package friend

import (
	"context"
	"im-chat/apps/social/rpc/socialclient"
	"im-chat/apps/user/rpc/userclient"
	"im-chat/pkg/ctxdata"

	"im-chat/apps/social/api/internal/svc"
	"im-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友列表
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// todo: add your logic here and delete this line

	uid := ctxdata.GetUid(l.ctx)

	friendList, err := l.svcCtx.FriendList(l.ctx, &socialclient.FriendListReq{
		UserId: uid,
	})

	if err != nil {
		return nil, err
	}

	if len(friendList.List) == 0 {
		return &types.FriendListResp{}, nil
	}

	// 根据好友id获取用户详情
	uids := make([]string, 0, len(friendList.List))
	for _, item := range friendList.List {
		uids = append(uids, item.FriendUid)
	}

	// 根据uids获取用户信息
	users, err := l.svcCtx.User.FindUser(l.ctx, &userclient.FindUserReq{
		Ids: uids,
	})

	if err != nil {
		return &types.FriendListResp{}, err
	}

	userRecords := make(map[string]*userclient.UserEntity, len(users.User))
	for i, _ := range users.User {
		userRecords[users.User[i].Id] = users.User[i]
	}

	respList := make([]*types.Friends, 0, len(friendList.List))
	for _, v := range friendList.List {
		friends := &types.Friends{
			Id:        v.Id,
			FriendUid: v.FriendUid,
		}

		if u, ok := userRecords[v.FriendUid]; ok {
			friends.Nickname = u.Nickname
			friends.Avatar = u.Avatar
		}

		respList = append(respList, friends)

	}

	return &types.FriendListResp{
		List: respList,
	}, nil
}
