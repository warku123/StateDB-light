package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSlotToAccessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSlotToAccessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSlotToAccessListLogic {
	return &AddSlotToAccessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSlotToAccessListLogic) AddSlotToAccessList(in *sdb.AddSlotToAccessListRequest) (*sdb.AddSlotToAccessListResponse, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.Statedb.AddSlotToAccessList(
		common.HexToAddress(in.Addr),
		common.HexToHash(in.Slot),
	)
	return &sdb.AddSlotToAccessListResponse{}, nil
}
