package mqclient

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"im-chat/apps/task/mq/mq"
)

type MsgChatTransferClient interface {
	Push(msg *mq.MsgChatTransfer) error
}

type msgChatTransferClient struct {
	pusher *kq.Pusher
}

func NewMsgChatTransferClient(addr []string, topic string,
	opts ...kq.PushOption) MsgChatTransferClient {
	return &msgChatTransferClient{pusher: kq.NewPusher(addr, topic, opts...)}
}

func (m *msgChatTransferClient) Push(msg *mq.MsgChatTransfer) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return m.pusher.Push(string(body))
}
