package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSlotToAccessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSlotToAccessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSlotToAccessListLogic {
	return &AddSlotToAccessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSlotToAccessListLogic) AddSlotToAccessList(req *types.AddSlotToAccessListRequest) (resp *types.AddSlotToAccessListResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.SdbRpc.AddSlotToAccessList(l.ctx, &sdb.AddSlotToAccessListRequest{
		Addr: req.Addr,
		Slot: req.Slot,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddSlotToAccessListResponse{}, nil
}
