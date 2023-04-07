package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetNonceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetNonceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetNonceLogic {
	return &SetNonceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetNonceLogic) SetNonce(req *types.SetNonceRequest) (resp *types.SetNonceResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.SetNonce(l.ctx, &sdb.SetNonceRequest{
		Addr:   req.Addr,
		Amount: req.Amount,
	})

	if err != nil {
		return nil, err
	}

	return &types.SetNonceResponse{
		Empty: res.Empty,
	}, nil
}
