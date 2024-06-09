package handler

import (
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
	})
}
