package user

import (
	"context"
	"errors"

	"mmt/mmt/internal/svc"
	"mmt/mmt/internal/types"
	"mmt/mmt/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateReq) (resp *types.NoneRes, err error) {
	u := &model.MmtUsers{}
	copier.Copy(u,req)
	err = l.svcCtx.Mysql.Debug().Table(model.MmtUsers{}.TableName()).Updates(u).Error
	if err != nil{
		return nil,errors.New("update user error! ")
	}
	return
}
