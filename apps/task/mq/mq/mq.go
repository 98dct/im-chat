package mq

import "im-chat/pkg/constants"

type MsgChatTransfer struct {
	ConversationId     string `json:"conversationId"`
	constants.ChatType `json:"chatType"`
	SendId             string   `json:"sendId"`
	RecvId             string   `json:"recvId"`
	RecvIds            []string `json:"recvIds"`
	SendTime           int64    `json:"sendTime"`

	constants.MType `json:"mType"`
	Content         string `json:"content"`
}

type MsgMarkRead struct {
	constants.ChatType `json:"chatType"`
	SendId             string   `json:"sendId"`
	RecvId             string   `json:"recvId"`
	ConversationId     string   `json:"conversationId"`
	MsgIds             []string `json:"msgIds"`
}
