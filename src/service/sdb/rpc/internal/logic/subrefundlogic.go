package logic

import (
	"context"

	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubRefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubRefundLogic {
	return &SubRefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubRefundLogic) SubRefund(in *sdb.SubRefundRequest) (*sdb.SubRefundResponse, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.Statedb.SubRefund(in.Amount)
	return &sdb.SubRefundResponse{}, nil
}
