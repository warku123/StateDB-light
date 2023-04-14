package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTransientStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTransientStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransientStateLogic {
	return &GetTransientStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTransientStateLogic) GetTransientState(in *sdb.GetTransientStateRequest) (*sdb.GetTransientStateResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.HexToAddress(in.Addr)
	real_key := common.HexToHash(in.Key)

	value := l.svcCtx.Statedb.GetTransientState(real_addr, real_key)

	return &sdb.GetTransientStateResponse{
		Value: value.String(),
	}, nil
}
