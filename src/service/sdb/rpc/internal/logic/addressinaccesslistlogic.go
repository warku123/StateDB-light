package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressInAccessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddressInAccessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressInAccessListLogic {
	return &AddressInAccessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddressInAccessListLogic) AddressInAccessList(in *sdb.AddressInAccessListRequest) (*sdb.AddressInAccessListResponse, error) {
	// todo: add your logic here and delete this line

	return &sdb.AddressInAccessListResponse{
		IsIn: l.svcCtx.Statedb.AddressInAccessList(
			common.BytesToAddress([]byte(in.Addr)),
		),
	}, nil
}
