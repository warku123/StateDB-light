package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeHashLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCodeHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeHashLogic {
	return &GetCodeHashLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCodeHashLogic) GetCodeHash(req *types.GetCodeHashRequest) (resp *types.GetCodeHashResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.GetCodeHash(l.ctx, &sdb.GetCodeHashRequest{
		Addr: req.Addr,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetCodeHashResponse{
		Hash: res.Hash,
	}, nil
}
