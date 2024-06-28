package logic

import (
	"context"
	"github.com/pkg/errors"
	"im-chat/pkg/xerr"

	"im-chat/apps/im/rpc/im"
	"im-chat/apps/im/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取会话记录
func (l *GetChatLogLogic) GetChatLog(in *im.GetChatLogReq) (*im.GetChatLogResp, error) {
	// todo: add your logic here and delete this line

	// 根据id
	if in.MsgId != "" {
		one, err := l.svcCtx.ChatLogModel.FindOne(l.ctx, in.MsgId)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "find chat log err: %v, msgId: ",
				err, in.MsgId)
		}
		return &im.GetChatLogResp{List: []*im.ChatLog{
			{
				Id:             one.ID.Hex(),
				ConversationId: one.ConversationId,
				SendId:         one.SendId,
				RecvId:         one.RecvId,
				MsgType:        int32(one.MsgType),
				MsgContent:     one.MsgContent,
				ChatType:       int32(one.ChatType),
				SendTime:       one.SendTime,
			},
		}}, nil
	}

	// 根据时间段分段查询
	data, err := l.svcCtx.ListBySendTime(l.ctx, in.ConversationId, in.StartSendTime, in.EndSendTime, in.Count)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "ListBySendTime err: %v, msgId: ",
			err, in.MsgId)
	}

	res := make([]*im.ChatLog, 0, len(data))
	for _, one := range data {
		res = append(res, &im.ChatLog{
			Id:             one.ID.Hex(),
			ConversationId: one.ConversationId,
			SendId:         one.SendId,
			RecvId:         one.RecvId,
			MsgType:        int32(one.MsgType),
			MsgContent:     one.MsgContent,
			ChatType:       int32(one.ChatType),
			SendTime:       one.SendTime,
		})
	}
	return &im.GetChatLogResp{
		List: res,
	}, nil
}
