package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeLogic {
	return &GetCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCodeLogic) GetCode(req *types.GetCodeRequest) (resp *types.GetCodeResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.SdbRpc.GetCode(l.ctx, &sdb.GetCodeRequest{
		Addr: req.Addr,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetCodeResponse{
		Code: common.Bytes2Hex(res.Code),
	}, nil
}
