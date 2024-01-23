package user

import (
	"context"
	"errors"

	"mmt/mmt/internal/svc"
	"mmt/mmt/internal/types"
	"mmt/mmt/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.GetUserByIdReq) (*types.GetUserByIdRes, error) {
	// todo: add your logic here and delete this line
	user := model.MmtUsers{}
	err := l.svcCtx.Mysql.Table(model.MmtUsers{}.TableName()).First(&user, "id", req.Id).Error
	if err != nil {
		return nil, errors.New("get user info error! ")
	}
	return &types.GetUserByIdRes{Data: user}, nil
}
