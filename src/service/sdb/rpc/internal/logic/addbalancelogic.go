package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBalanceLogic {
	return &AddBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBalanceLogic) AddBalance(in *sdb.AddBalanceRequest) (*sdb.AddBalanceResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.HexToAddress(in.Addr)
	amount, _ := common.StringToBig(in.Amount)
	l.svcCtx.Statedb.AddBalance(real_addr, amount)

	return &sdb.AddBalanceResponse{Empty: true}, nil
}
