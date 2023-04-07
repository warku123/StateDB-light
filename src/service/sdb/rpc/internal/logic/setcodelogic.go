package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCodeLogic {
	return &SetCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetCodeLogic) SetCode(in *sdb.SetCodeRequest) (*sdb.SetCodeResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.BytesToAddress([]byte(in.Addr))
	l.svcCtx.Statedb.SetCode(real_addr, []byte(in.Code))
	return &sdb.SetCodeResponse{}, nil
}
