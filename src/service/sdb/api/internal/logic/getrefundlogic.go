package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundLogic {
	return &GetRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRefundLogic) GetRefund(req *types.GetRefundRequest) (resp *types.GetRefundResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.GetRefund(l.ctx, &sdb.GetRefundRequest{})

	if err != nil {
		return nil, err
	}

	return &types.GetRefundResponse{
		Amount: res.Amount,
	}, nil
}
