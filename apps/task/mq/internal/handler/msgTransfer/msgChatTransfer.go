package msgTransfer

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"im-chat/apps/task/mq/internal/svc"
)

type MsgChatTransfer struct {
	logx.Logger
	svc *svc.ServiceContext
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		Logger: logx.WithContext(context.Background()),
		svc:    svc,
	}
}

func (m *MsgChatTransfer) Consume(key, value string) error {
	fmt.Println("key:", key, "value:", value)
	return nil
}
