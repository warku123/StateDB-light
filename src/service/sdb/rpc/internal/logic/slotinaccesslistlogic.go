package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SlotInAccessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSlotInAccessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SlotInAccessListLogic {
	return &SlotInAccessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SlotInAccessListLogic) SlotInAccessList(in *sdb.SlotInAccessListRequest) (*sdb.SlotInAccessListResponse, error) {
	// todo: add your logic here and delete this line
	addr_ok, slot_ok := l.svcCtx.Statedb.SlotInAccessList(
		common.HexToAddress(in.Addr),
		common.HexToHash(in.Hash),
	)
	return &sdb.SlotInAccessListResponse{
		AddrOk: addr_ok,
		SlotOk: slot_ok,
	}, nil
}
