package logic

import (
	"context"
	"math/big"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubBalanceLogic {
	return &SubBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubBalanceLogic) SubBalance(in *sdb.SubBalanceRequest) (*sdb.SubBalanceResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.BytesToAddress([]byte(in.Addr))
	l.svcCtx.Config.StateDB.SubBalance(real_addr, big.NewInt(in.Amount))

	return &sdb.SubBalanceResponse{Empty: true}, nil
}
