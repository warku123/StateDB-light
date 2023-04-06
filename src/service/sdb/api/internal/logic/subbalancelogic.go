package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubBalanceLogic {
	return &SubBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubBalanceLogic) SubBalance(req *types.SubBalanceRequest) (resp *types.SubBalanceResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.SubBalance(l.ctx, &sdb.SubBalanceRequest{
		Addr:   req.Addr,
		Amount: req.Amount,
	})

	if err != nil {
		return nil, err
	}

	return &types.SubBalanceResponse{
		Empty: res.Empty,
	}, nil
}
