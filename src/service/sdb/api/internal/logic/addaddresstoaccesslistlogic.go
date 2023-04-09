package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddressToAccessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAddressToAccessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressToAccessListLogic {
	return &AddAddressToAccessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAddressToAccessListLogic) AddAddressToAccessList(req *types.AddAddressToAccessListRequest) (resp *types.AddAddressToAccessListResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.SdbRpc.AddAddressToAccessList(l.ctx, &sdb.AddAddressToAccessListRequest{
		Addr: req.Addr,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddAddressToAccessListResponse{}, nil
}
