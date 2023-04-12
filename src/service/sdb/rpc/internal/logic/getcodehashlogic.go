package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCodeHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeHashLogic {
	return &GetCodeHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCodeHashLogic) GetCodeHash(in *sdb.GetCodeHashRequest) (*sdb.GetCodeHashResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.BytesToAddress([]byte(in.Addr))
	res := l.svcCtx.Statedb.GetCodeHash(real_addr)
	return &sdb.GetCodeHashResponse{
		Hash: res.Hex(),
	}, nil
}
