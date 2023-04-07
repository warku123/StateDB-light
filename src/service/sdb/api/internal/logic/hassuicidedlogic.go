package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HasSuicidedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHasSuicidedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HasSuicidedLogic {
	return &HasSuicidedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HasSuicidedLogic) HasSuicided(req *types.HasSuicidedRequest) (resp *types.HasSuicidedResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.HasSuicided(l.ctx, &sdb.HasSuicidedRequest{
		Addr: req.Addr,
	})

	if err != nil {
		return nil, err
	}

	return &types.HasSuicidedResponse{
		Is_suicide: res.IsSuicide,
	}, nil
}
