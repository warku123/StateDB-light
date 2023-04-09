package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RevertToSnapshotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRevertToSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevertToSnapshotLogic {
	return &RevertToSnapshotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RevertToSnapshotLogic) RevertToSnapshot(req *types.RevertToSnapshotRequest) (resp *types.RevertToSnapshotResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.SdbRpc.RevertToSnapshot(l.ctx, &sdb.RevertToSnapshotRequest{
		Revid: req.Revid,
	})

	if err != nil {
		return nil, err
	}

	return &types.RevertToSnapshotResponse{}, nil
}
