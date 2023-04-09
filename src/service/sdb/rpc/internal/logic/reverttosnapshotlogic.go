package logic

import (
	"context"

	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RevertToSnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRevertToSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevertToSnapshotLogic {
	return &RevertToSnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RevertToSnapshotLogic) RevertToSnapshot(in *sdb.RevertToSnapshotRequest) (*sdb.RevertToSnapshotResponse, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.Statedb.RevertToSnapshot(int(in.Revid))
	return &sdb.RevertToSnapshotResponse{}, nil
}
