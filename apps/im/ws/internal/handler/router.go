package handler

import (
	"im-chat/apps/im/ws/internal/handler/conversation"
	"im-chat/apps/im/ws/internal/handler/push"
	"im-chat/apps/im/ws/internal/handler/user"
	"im-chat/apps/im/ws/internal/svc"
	"im-chat/apps/im/ws/websocket"
)

func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoutes([]websocket.Route{
		{
			Method:  "user.online",
			Handler: user.OnLine(svc),
		},
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "push.push",
			Handler: push.Push(svc),
		},
	})
}
