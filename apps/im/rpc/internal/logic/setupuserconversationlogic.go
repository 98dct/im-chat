package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"im-chat/apps/im/immodels"
	"im-chat/apps/im/rpc/im"
	"im-chat/apps/im/rpc/internal/svc"
	"im-chat/pkg/constants"
	"im-chat/pkg/wuid"
	"im-chat/pkg/xerr"
)

type SetUpUserConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUpUserConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUpUserConversationLogic {
	return &SetUpUserConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 建立会话：群聊、私聊
func (l *SetUpUserConversationLogic) SetUpUserConversation(in *im.SetUpUserConversationReq) (*im.SetUpUserConversationResp, error) {
	// todo: add your logic here and delete this line

	switch constants.ChatType(in.ChatType) {
	case constants.SingleChatType:

		// 生成会话的id
		conversationId := wuid.CombineId(in.SendId, in.RecvId)
		// 验证是否建立过会话
		one, err := l.svcCtx.ConversationModel.FindOne(l.ctx, conversationId)
		if err != nil {
			if err == immodels.ErrNotFound {
				err := l.svcCtx.ConversationModel.Insert(l.ctx, &immodels.Conversation{
					ConversationId: conversationId,
					ChatType:       constants.SingleChatType,
				})
				if err != nil {
					return nil, errors.Wrapf(xerr.NewDBErr(), "ConversationModel.Insert err: %v, conversationId: ",
						err, conversationId)
				}
			} else {
				return nil, errors.Wrapf(xerr.NewDBErr(), "ConversationModel.FindOne err: %v, conversationId: ",
					err, conversationId)
			}

		} else if one != nil {
			return nil, nil
		}
		// 建立两者的会话
		err = l.setUpUserConversation(conversationId, in.SendId, in.RecvId, constants.SingleChatType, true)
		if err != nil {
			return nil, err
		}
		err = l.setUpUserConversation(conversationId, in.RecvId, in.SendId, constants.SingleChatType, false)
		if err != nil {
			return nil, err
		}

	}

	return &im.SetUpUserConversationResp{}, nil
}

func (l *SetUpUserConversationLogic) setUpUserConversation(conversationId, userId, recvId string, chatType constants.ChatType,
	isShow bool) error {

	// 用户的会话列表
	conversations, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, userId)
	if err != nil {
		if err == immodels.ErrNotFound {
			conversations = &immodels.Conversations{
				ID:               primitive.NewObjectID(),
				UserId:           userId,
				ConversationList: make(map[string]*immodels.Conversation),
			}
		} else {
			return errors.Wrapf(xerr.NewDBErr(), "ConversationsModel.FindByUserId err: %v, userId: ",
				err, userId)
		}
	}

	// 更新会话记录
	if _, ok := conversations.ConversationList[conversationId]; ok {
		return nil
	}

	conversations.ConversationList[conversationId] = &immodels.Conversation{
		ConversationId: conversationId,
		ChatType:       constants.SingleChatType,
		IsShow:         isShow,
	}

	_, err = l.svcCtx.ConversationsModel.Update(l.ctx, conversations)
	if err != nil {
		return errors.Wrapf(xerr.NewDBErr(), "ConversationsModel.Update err: %v, conversations: ",
			err, conversations)
	}

	return nil
}
