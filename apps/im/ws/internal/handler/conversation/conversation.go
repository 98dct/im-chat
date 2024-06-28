package conversation

import (
	"github.com/mitchellh/mapstructure"
	"im-chat/apps/im/ws/internal/svc"
	"im-chat/apps/im/ws/websocket"
	"im-chat/apps/im/ws/ws"
	"im-chat/apps/task/mq/mq"
	"im-chat/pkg/constants"
	"im-chat/pkg/wuid"
	"time"
)

// 私聊功能
// 1.消息类型
// 2.会话id
// 3.消息记录
// 4.查找并发送消息给目标用户

func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {

		// 私聊
		var data ws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		if data.ConversationId == "" {
			data.ConversationId = wuid.CombineId(data.SendId, data.RecvId)
		}

		switch data.ChatType {
		case constants.SingleChatType:

			err := svc.MsgChatTransferClient.Push(&mq.MsgChatTransfer{
				ConversationId: data.ConversationId,
				ChatType:       data.ChatType,
				SendId:         conn.Uid,
				RecvId:         data.RecvId,
				SendTime:       time.Now().UnixNano(),
				MType:          data.Msg.MType,
				Content:        data.Msg.Content,
			})

			if err != nil {
				srv.Send(websocket.NewErrMessage(err), conn)
				return
			}

			//err := logic.NewConversation(context.Background(), srv, svc).SingleChat(&data, conn.Uid)
			//if err != nil {
			//	srv.Send(websocket.NewErrMessage(err), conn)
			//	return
			//}
			//
			//srv.SendByIds(websocket.NewMessage(conn.Uid, ws.Chat{
			//	ConversationId: data.ConversationId,
			//	ChatType:       data.ChatType,
			//	SendId:         conn.Uid,
			//	RecvId:         data.RecvId,
			//	Msg:            data.Msg,
			//	SendTime:       time.Now().UnixMilli(),
			//}), data.RecvId)

		}

	}
}
