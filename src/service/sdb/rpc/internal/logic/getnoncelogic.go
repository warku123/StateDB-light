package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNonceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNonceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNonceLogic {
	return &GetNonceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNonceLogic) GetNonce(in *sdb.GetNonceRequest) (*sdb.GetNonceResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.BytesToAddress([]byte(in.Addr))
	nonce := l.svcCtx.Config.StateDB.GetNonce(real_addr)

	return &sdb.GetNonceResponse{
		Amount: nonce,
	}, nil
}
