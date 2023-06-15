package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPreimageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPreimageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPreimageLogic {
	return &AddPreimageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPreimageLogic) AddPreimage(req *types.AddPreimageRequest) (resp *types.AddPreimageResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.SdbRpc.AddPreimage(l.ctx, &sdb.AddPreimageRequest{
		Hash:     req.Hash,
		Preimage: common.Hex2Bytes(req.Preimage),
	})
	if err != nil {
		return nil, err
	}

	return &types.AddPreimageResponse{}, nil
}
