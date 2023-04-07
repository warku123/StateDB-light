package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmptyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmptyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmptyLogic {
	return &EmptyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmptyLogic) Empty(req *types.EmptyRequest) (resp *types.EmptyResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.Empty(l.ctx, &sdb.EmptyRequest{
		Addr: req.Addr,
	})

	if err != nil {
		return nil, err
	}

	return &types.EmptyResponse{
		Is_empty: res.Is_Empty,
	}, nil
}
