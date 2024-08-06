package push

import (
	"github.com/mitchellh/mapstructure"
	"im-chat/apps/im/ws/internal/svc"
	"im-chat/apps/im/ws/websocket"
	"im-chat/apps/im/ws/ws"
	"im-chat/pkg/constants"
)

func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data ws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err))
			return
		}

		// 要发送的目标

		switch data.ChatType {
		case constants.SingleChatType:
			Single(srv, &data, data.RecvId)
		case constants.GroupChatType:
			Group(srv, &data)
		}

	}
}

func Single(srv *websocket.Server, data *ws.Push, recvId string) error {

	rconn := srv.GetConn(recvId)

	if rconn == nil {
		// todo 用户离线
		return nil
	}

	return srv.Send(websocket.NewMessage(data.SendId, &ws.Chat{
		ConversationId: data.ConversationId,
		ChatType:       data.ChatType,
		Msg: ws.Msg{
			MsgId:       data.MsgId,
			ReadRecords: data.ReadRecords,
			MType:       data.MType,
			Content:     data.Content,
		},
		SendTime: data.SendTime,
	}), rconn)
}

func Group(srv *websocket.Server, data *ws.Push) error {
	for _, id := range data.RecvIds {
		func(id string) {
			srv.Schedule(func() {
				Single(srv, data, id)
			})
		}(id)
	}

	return nil
}
