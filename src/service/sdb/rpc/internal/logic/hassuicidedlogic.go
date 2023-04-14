package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HasSuicidedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHasSuicidedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HasSuicidedLogic {
	return &HasSuicidedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HasSuicidedLogic) HasSuicided(in *sdb.HasSuicidedRequest) (*sdb.HasSuicidedResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.HexToAddress(in.Addr)
	is_suicide := l.svcCtx.Statedb.HasSuicided(real_addr)

	return &sdb.HasSuicidedResponse{IsSuicide: is_suicide}, nil
}
