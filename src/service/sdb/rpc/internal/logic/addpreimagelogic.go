package logic

import (
	"context"

	"statedbl/common"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPreimageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPreimageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPreimageLogic {
	return &AddPreimageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPreimageLogic) AddPreimage(in *sdb.AddPreimageRequest) (*sdb.AddPreimageResponse, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.Statedb.AddPreimage(
		common.HexToHash(in.Hash),
		[]byte(in.Preimage),
	)
	return &sdb.AddPreimageResponse{}, nil
}
