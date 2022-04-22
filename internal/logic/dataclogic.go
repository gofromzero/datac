package logic

import (
	"context"

	"datac/internal/svc"
	"datac/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DatacLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDatacLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DatacLogic {

	return &DatacLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DatacLogic) Datac(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Error(req)

	return
}
