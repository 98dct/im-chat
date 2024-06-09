package immodels

import (
	"im-chat/pkg/constants"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatLog struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ConversationId string             `bson:"conversationId"`
	SendId         string             `bson:"sendId"`
	RecvId         string             `bson:"recvId"`
	ChatType       constants.ChatType `bson:"chatType"`
	MsgFrom        int                `bson:"msgFrom"`
	MsgType        constants.MType    `bson:"msgType"`
	MsgContent     string             `bson:"msgContent"`
	SendTime       int64              `bson:"sendTime"`
	Status         int                `bson:"status"`

	UpdateAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	CreateAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}
