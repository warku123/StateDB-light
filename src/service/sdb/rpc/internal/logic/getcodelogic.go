package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeLogic {
	return &GetCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCodeLogic) GetCode(in *sdb.GetCodeRequest) (*sdb.GetCodeResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.BytesToAddress([]byte(in.Addr))
	res := l.svcCtx.Statedb.GetCode(real_addr)
	return &sdb.GetCodeResponse{
		Code: string(res),
	}, nil
}
