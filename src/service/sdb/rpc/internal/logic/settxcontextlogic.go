package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTxContextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetTxContextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTxContextLogic {
	return &SetTxContextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetTxContextLogic) SetTxContext(in *sdb.SetTxContextRequest) (*sdb.SetTxContextResponse, error) {
	real_tx_hash := common.HexToHash(in.TxHash)
	l.svcCtx.Statedb.SetTxContext(real_tx_hash, int(in.Index))

	return &sdb.SetTxContextResponse{}, nil
}
