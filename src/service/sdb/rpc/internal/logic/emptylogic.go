package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmptyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmptyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmptyLogic {
	return &EmptyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmptyLogic) Empty(in *sdb.EmptyRequest) (*sdb.EmptyResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.BytesToAddress([]byte(in.Addr))

	return &sdb.EmptyResponse{
		Is_Empty: l.svcCtx.Statedb.Empty(real_addr),
	}, nil
}
