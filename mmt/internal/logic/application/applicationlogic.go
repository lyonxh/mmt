package application

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mmt/mmt/internal/svc"
)

type ApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplicationLogic {
	return &ApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplicationLogic) Application() error {
	// todo: add your logic here and delete this line

	return nil
}
