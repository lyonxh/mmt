package user

import (
	"context"
	"errors"
	"time"

	"mmt/common"
	"mmt/mmt/internal/svc"
	"mmt/mmt/internal/types"
	"mmt/mmt/model"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginRes, error) {
	user := model.MmtUsers{}
	if err := l.svcCtx.Mysql.Table(user.TableName()).Where("user_name = ?", req.UserName).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("search user info error! ")
	}
	if user.UserName == "" {
		return nil, errors.New("user not exist! ")
	}
	if req.PassWord != user.PassWord {
		return nil, errors.New("password error! ")
	}
	now := time.Now().Unix()
	accessToken, err := common.CreateJwtToken(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire, now, user)
	if err != nil {
		return nil, errors.New("generate token fail! ")
	}
	return &types.LoginRes{
		Info:         user,
		AccessToken:  accessToken,
		AccessExpire: now + l.svcCtx.Config.Auth.AccessExpire,
		// RefreshAfter: (now + l.svcCtx.Config.Auth.AccessExpire) - 5,
		RefreshAfter: (now + l.svcCtx.Config.Auth.AccessExpire) - 3600, //	提前1h刷新
	}, nil
}
