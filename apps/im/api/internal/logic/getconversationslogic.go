package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"im-chat/apps/im/rpc/imclient"
	"im-chat/pkg/ctxdata"

	"im-chat/apps/im/api/internal/svc"
	"im-chat/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取会话
func NewGetConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConversationsLogic {
	return &GetConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConversationsLogic) GetConversations(req *types.GetConversationsReq) (resp *types.GetConversationsResp, err error) {
	// todo: add your logic here and delete this line

	uid := ctxdata.GetUid(l.ctx)
	conversations, err := l.svcCtx.Im.GetConversations(l.ctx, &imclient.GetConversationsReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	var res types.GetConversationsResp
	copier.Copy(&res, conversations)

	return &res, nil
}
