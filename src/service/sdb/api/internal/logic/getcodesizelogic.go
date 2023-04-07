package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeSizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCodeSizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeSizeLogic {
	return &GetCodeSizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCodeSizeLogic) GetCodeSize(req *types.GetCodeSizeRequest) (resp *types.GetCodeSizeResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.GetCodeSize(l.ctx, &sdb.GetCodeSizeRequest{
		Addr: req.Addr,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetCodeSizeResponse{
		Size: res.Size,
	}, nil
}
