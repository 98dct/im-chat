package msgTransfer

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"im-chat/apps/im/ws/websocket"
	"im-chat/apps/im/ws/ws"
	"im-chat/apps/social/rpc/socialclient"
	"im-chat/apps/task/mq/internal/svc"
	"im-chat/pkg/constants"
)

type baseMsgTransfer struct {
	svc *svc.ServiceContext
	logx.Logger
}

func NewBaseMsgTransfer(svc *svc.ServiceContext) *baseMsgTransfer {
	return &baseMsgTransfer{
		svc:    svc,
		Logger: logx.WithContext(context.Background()),
	}
}

func (b *baseMsgTransfer) Transfer(ctx context.Context, data *ws.Push) error {

	var err error
	switch data.ChatType {
	case constants.SingleChatType:
		err = b.single(ctx, data)
	case constants.GroupChatType:
		err = b.group(ctx, data)
	}

	return err
}

func (m *baseMsgTransfer) single(ctx context.Context, data *ws.Push) error {
	return m.svc.WsClient.Send(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push.push",
		FormId:    constants.SYSTEM_ROOT_UID,
		Data:      data,
	})
}

func (m *baseMsgTransfer) group(ctx context.Context, data *ws.Push) error {
	groupUsers, err := m.svc.Social.GroupUsers(ctx, &socialclient.GroupUsersReq{
		GroupId: data.RecvId,
	})
	if err != nil {
		return err
	}

	data.RecvIds = make([]string, 0, len(groupUsers.List))

	for _, members := range groupUsers.List {
		if members.UserId == data.SendId {
			continue
		}

		data.RecvIds = append(data.RecvIds, members.UserId)
	}

	return m.svc.WsClient.Send(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push.push",
		FormId:    constants.SYSTEM_ROOT_UID,
		Data:      data,
	})
}
