package user

import (
	"context"
	"errors"

	"mmt/mmt/internal/svc"
	"mmt/mmt/internal/types"
	"mmt/mmt/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (resp *types.NoneRes, err error) {
	err = l.svcCtx.Mysql.Table(model.MmtUsers{}.TableName()).Delete(&model.MmtUsers{}, req.Id).Error
	if err != nil {
		return nil, errors.New("delete user error! ")
	}
	return
}
