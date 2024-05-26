package logic

import (
	"context"
	"github.com/pkg/errors"
	"im-chat/apps/user/models"
	"im-chat/pkg/ctxdata"
	"im-chat/pkg/encrypt"
	"im-chat/pkg/xerr"
	"time"

	"im-chat/apps/user/rpc/internal/svc"
	"im-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrNotRegister = xerr.New(xerr.SERVER_COMMON_ERROR, "手机号没注册！") //errors.New("用户尚未注册！")
	ErrPassword    = xerr.New(xerr.SERVER_COMMON_ERROR, "密码错误！")
	ErrGetToken    = xerr.New(xerr.SERVER_COMMON_ERROR, "获取token出错！")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line

	// 1.校验是否已经注册
	userEntity, err := l.svcCtx.UsersModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, errors.WithStack(ErrNotRegister)
		}
		return nil, errors.Wrapf(ErrNotRegister, "find user by phone err %v,req %v",
			err, in.Phone)
	}

	//2.密码校验
	if !encrypt.ValidatePasswordHash(in.Password, userEntity.Password.String) {
		return nil, errors.WithStack(ErrPassword)
	}

	//3.生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now,
		l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)
	if err != nil {
		return nil, errors.Wrapf(ErrGetToken, "get jwt token err %v", err)
	}

	return &user.LoginResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
