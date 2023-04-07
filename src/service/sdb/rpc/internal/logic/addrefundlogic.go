package logic

import (
	"context"

	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRefundLogic {
	return &AddRefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRefundLogic) AddRefund(in *sdb.AddRefundRequest) (*sdb.AddRefundResponse, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.Statedb.AddRefund(in.Amount)
	return &sdb.AddRefundResponse{}, nil
}
