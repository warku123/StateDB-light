package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTransientStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetTransientStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTransientStateLogic {
	return &SetTransientStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetTransientStateLogic) SetTransientState(req *types.SetTransientStateRequest) (resp *types.SetTransientStateResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.SdbRpc.SetTransientState(l.ctx, &sdb.SetTransientStateRequest{
		Addr:  req.Addr,
		Key:   req.Key,
		Value: req.Value,
	})

	if err != nil {
		return nil, err
	}

	return &types.SetTransientStateResponse{}, nil
}
