package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRefundLogic {
	return &AddRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRefundLogic) AddRefund(req *types.AddRefundRequest) (resp *types.AddRefundResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.SdbRpc.AddRefund(l.ctx, &sdb.AddRefundRequest{
		Amount: req.Amount,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddRefundResponse{}, nil
}
