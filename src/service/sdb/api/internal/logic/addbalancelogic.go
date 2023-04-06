package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBalanceLogic {
	return &AddBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBalanceLogic) AddBalance(req *types.AddBalanceRequest) (resp *types.AddBalanceResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.AddBalance(l.ctx, &sdb.AddBalanceRequest{
		Addr:   req.Addr,
		Amount: req.Amount,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddBalanceResponse{
		Empty: res.Empty,
	}, nil
}
