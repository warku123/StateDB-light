package logic

import (
	"context"

	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundLogic {
	return &GetRefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRefundLogic) GetRefund(in *sdb.GetRefundRequest) (*sdb.GetRefundResponse, error) {
	// todo: add your logic here and delete this line

	return &sdb.GetRefundResponse{
		Amount: l.svcCtx.Statedb.GetRefund(),
	}, nil
}
