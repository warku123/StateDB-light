package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistLogic {
	return &ExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExistLogic) Exist(in *sdb.ExistRequest) (*sdb.ExistResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.HexToAddress(in.Addr)

	return &sdb.ExistResponse{
		Is_Exist: l.svcCtx.Statedb.Exist(real_addr),
	}, nil
}
