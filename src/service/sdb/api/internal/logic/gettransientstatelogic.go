package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTransientStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTransientStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransientStateLogic {
	return &GetTransientStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTransientStateLogic) GetTransientState(req *types.GetTransientStateRequest) (resp *types.GetTransientStateResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.GetTransientState(l.ctx, &sdb.GetTransientStateRequest{
		Addr: req.Addr,
		Key:  req.Key,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetTransientStateResponse{
		Value: res.Value,
	}, nil
}
