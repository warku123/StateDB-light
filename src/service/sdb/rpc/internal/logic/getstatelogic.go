package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStateLogic {
	return &GetStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStateLogic) GetState(in *sdb.GetStateRequest) (*sdb.GetStateResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.HexToAddress(in.Addr)
	real_hash := common.HexToHash(in.Hash)
	res := l.svcCtx.Statedb.GetState(real_addr, real_hash)
	return &sdb.GetStateResponse{
		Hash: res.Hex(),
	}, nil
}
