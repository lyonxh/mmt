package user

import (
	"context"
	"errors"

	"mmt/mmt/internal/svc"
	"mmt/mmt/internal/types"
	"mmt/mmt/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (*types.GetUserListRes, error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := l.svcCtx.Mysql.Table(model.MmtUsers{}.TableName())

	if req.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+req.UserName+"%")
	}
	if req.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+req.NickName+"%")
	}
	if req.Mobile != "" {
		db = db.Where("mobile LIKE ?", "%"+req.Mobile+"%")
	}
	if req.IdCard != "" {
		db = db.Where("id_card LIKE ?", "%"+req.IdCard+"%")
	}
	if req.Email != "" {
		db = db.Where("email LIKE ?", "%"+req.Email+"%")
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, errors.New("count error! ")
	}
	user := []model.MmtUsers{}
	if err := db.Find(&user).Order("id desc").Limit(limit).Offset(offset).Error; err != nil {
		return nil, errors.New("search user list error! ")
	}

	return &types.GetUserListRes{
		Info: user,
		Pagination: types.Pagination{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}
