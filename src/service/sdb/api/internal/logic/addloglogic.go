package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogLogic {
	return &AddLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogLogic) AddLog(req *types.AddLogRequest) (resp *types.AddLogResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.SdbRpc.AddLog(l.ctx, &sdb.AddLogRequest{
		Addr:        req.Address,
		Topic:       req.Topics,
		Data:        req.Data,
		BlockNumber: req.BlockNumber,
		TxHash:      req.TxHash,
		TxIndex:     uint32(req.TxIndex),
		BlockHash:   req.BlockHash,
		Index:       req.Index,
		Removed:     req.Removed,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddLogResponse{}, nil
}
