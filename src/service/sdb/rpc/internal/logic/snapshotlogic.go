package logic

import (
	"context"

	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SnapshotLogic {
	return &SnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SnapshotLogic) Snapshot(in *sdb.SnapshotRequest) (*sdb.SnapshotResponse, error) {
	// todo: add your logic here and delete this line

	return &sdb.SnapshotResponse{
		Revid: int32(l.svcCtx.Statedb.Snapshot()),
	}, nil
}
