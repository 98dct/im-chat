package push

import (
	"github.com/mitchellh/mapstructure"
	"im-chat/apps/im/ws/internal/svc"
	"im-chat/apps/im/ws/websocket"
	"im-chat/apps/im/ws/ws"
)

func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data ws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err))
			return
		}

		// 要发送的目标
		rconn := srv.GetConn(data.RecvId)

		if rconn == nil {
			// todo 用户离线
			return
		}

		srv.Send(websocket.NewMessage(data.SendId, &ws.Chat{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			Msg: ws.Msg{
				MType:   data.MType,
				Content: data.Content,
			},
			SendTime: data.SendTime,
		}), rconn)

	}
}
