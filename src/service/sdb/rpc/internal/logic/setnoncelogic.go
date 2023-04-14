package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetNonceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetNonceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetNonceLogic {
	return &SetNonceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetNonceLogic) SetNonce(in *sdb.SetNonceRequest) (*sdb.SetNonceResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.HexToAddress(in.Addr)
	l.svcCtx.Statedb.SetNonce(real_addr, in.Amount)

	return &sdb.SetNonceResponse{
		Empty: true,
	}, nil
}
