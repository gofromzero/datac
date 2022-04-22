package logic

import (
	"context"

	"datac/internal/svc"
	"datac/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DatacHandler2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDatacHandler2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *DatacHandler2Logic {
	return &DatacHandler2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DatacHandler2Logic) DatacHandler2(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
