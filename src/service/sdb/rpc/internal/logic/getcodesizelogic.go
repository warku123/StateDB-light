package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeSizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCodeSizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeSizeLogic {
	return &GetCodeSizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCodeSizeLogic) GetCodeSize(in *sdb.GetCodeSizeRequest) (*sdb.GetCodeSizeResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.BytesToAddress([]byte(in.Addr))
	codesize := l.svcCtx.Statedb.GetCodeSize(real_addr)
	return &sdb.GetCodeSizeResponse{
		Size: int32(codesize),
	}, nil
}
