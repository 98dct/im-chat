package ws

import "im-chat/pkg/constants"

type Msg struct {
	constants.MType `mapstructure:"mType"`
	Content         string `mapstructure:"content"`
}

type Chat struct {
	ConversationId     string `mapstructure:"conversationId"`
	constants.ChatType `mapstructure:"chatType"`
	SendId             string `mapstructure:"sendId"`
	RecvId             string `mapstructure:"recvId"`
	Msg                `mapstructure:"msg"`
	SendTime           int64 `mapstructure:"sendTime"`
}

type Push struct {
	ConversationId     string `mapstructure:"conversationId"`
	constants.ChatType `mapstructure:"chatType"`
	SendId             string `mapstructure:"sendId"`
	RecvId             string `mapstructure:"recvId"`
	SendTime           int64  `mapstructure:"sendTime"`

	constants.MType `mapstructure:"mType"`
	Content         string `mapstructure:"content"`
}
