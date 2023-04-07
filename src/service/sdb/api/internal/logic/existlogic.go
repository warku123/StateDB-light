package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistLogic {
	return &ExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExistLogic) Exist(req *types.ExistRequest) (resp *types.ExistResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.Exist(l.ctx, &sdb.ExistRequest{
		Addr: req.Addr,
	})

	if err != nil {
		return nil, err
	}

	return &types.ExistResponse{
		Is_exist: res.Is_Exist,
	}, nil
}
