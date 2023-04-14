package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetStateLogic {
	return &SetStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetStateLogic) SetState(in *sdb.SetStateRequest) (*sdb.SetStateResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.HexToAddress(in.Addr)
	real_key := common.HexToHash(in.Key)
	real_value := common.HexToHash(in.Value)
	l.svcCtx.Statedb.SetState(real_addr, real_key, real_value)
	return &sdb.SetStateResponse{}, nil
}
