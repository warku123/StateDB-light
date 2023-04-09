package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SlotInAccessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSlotInAccessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SlotInAccessListLogic {
	return &SlotInAccessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SlotInAccessListLogic) SlotInAccessList(req *types.SlotInAccessListRequest) (resp *types.SlotInAccessListResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.SlotInAccessList(l.ctx, &sdb.SlotInAccessListRequest{
		Addr: req.Addr,
		Hash: req.Hash,
	})

	if err != nil {
		return nil, err
	}

	return &types.SlotInAccessListResponse{
		Addr_ok: res.AddrOk,
		Slot_ok: res.SlotOk,
	}, nil
}
