package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommittedStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommittedStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommittedStateLogic {
	return &GetCommittedStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommittedStateLogic) GetCommittedState(in *sdb.GetCommittedStateRequest) (*sdb.GetCommittedStateResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.HexToAddress(in.Addr)
	real_hash := common.HexToHash(in.Hash)
	res := l.svcCtx.Statedb.GetCommittedState(real_addr, real_hash)
	return &sdb.GetCommittedStateResponse{
		Hash: res.Hex(),
	}, nil
}
