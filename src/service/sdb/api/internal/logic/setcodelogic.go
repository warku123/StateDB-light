package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetCodeLogic {
	return &SetCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetCodeLogic) SetCode(req *types.SetCodeRequest) (resp *types.SetCodeResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.SdbRpc.SetCode(l.ctx, &sdb.SetCodeRequest{
		Addr: req.Addr,
		Code: req.Code,
	})

	if err != nil {
		return nil, err
	}

	return &types.SetCodeResponse{}, nil
}
