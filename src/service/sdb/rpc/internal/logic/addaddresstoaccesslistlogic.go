package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddressToAccessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAddressToAccessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressToAccessListLogic {
	return &AddAddressToAccessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAddressToAccessListLogic) AddAddressToAccessList(in *sdb.AddAddressToAccessListRequest) (*sdb.AddAddressToAccessListResponse, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.Statedb.AddAddressToAccessList(common.BytesToAddress([]byte(in.Addr)))
	return &sdb.AddAddressToAccessListResponse{}, nil
}
