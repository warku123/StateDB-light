package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressInAccessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressInAccessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressInAccessListLogic {
	return &AddressInAccessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressInAccessListLogic) AddressInAccessList(req *types.AddressInAccessListRequest) (resp *types.AddressInAccessListResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.AddressInAccessList(l.ctx, &sdb.AddressInAccessListRequest{
		Addr: req.Addr,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddressInAccessListResponse{
		Is_in: res.IsIn,
	}, nil
}
