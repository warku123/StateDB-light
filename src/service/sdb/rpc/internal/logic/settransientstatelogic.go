package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetTransientStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetTransientStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetTransientStateLogic {
	return &SetTransientStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetTransientStateLogic) SetTransientState(in *sdb.SetTransientStateRequest) (*sdb.SetTransientStateResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.HexToAddress(in.Addr)
	real_key := common.HexToHash(in.Key)
	real_value := common.HexToHash(in.Value)

	l.svcCtx.Statedb.SetTransientState(real_addr, real_key, real_value)

	return &sdb.SetTransientStateResponse{}, nil
}
