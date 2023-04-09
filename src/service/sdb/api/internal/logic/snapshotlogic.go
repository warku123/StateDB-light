package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SnapshotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SnapshotLogic {
	return &SnapshotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SnapshotLogic) Snapshot(req *types.SnapshotRequest) (resp *types.SnapshotResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.Snapshot(l.ctx, &sdb.SnapshotRequest{})

	if err != nil {
		return nil, err
	}

	return &types.SnapshotResponse{
		Revid: res.Revid,
	}, nil
}
