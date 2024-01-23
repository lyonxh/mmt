package user

import (
	"context"
	"errors"

	"mmt/mmt/internal/svc"
	"mmt/mmt/internal/types"
	"mmt/mmt/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterRes, error) {
	u := &model.MmtUsers{}
	_ = copier.Copy(u, req)
	if !errors.Is(l.svcCtx.Mysql.Where("user_name = ?", req.UserName).First(&model.MmtUsers{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user has exist")
	}
	if err := l.svcCtx.Mysql.Table(model.MmtUsers{}.TableName()).Create(u).Error; err != nil {
		return nil, err
	}
	return &types.RegisterRes{RegisterReq: *req}, nil
}
