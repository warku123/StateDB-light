package logic

import (
	"context"

	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"statedbl/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type SuicideLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSuicideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuicideLogic {
	return &SuicideLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SuicideLogic) Suicide(in *sdb.SuicideRequest) (*sdb.SuicideResponse, error) {
	// todo: add your logic here and delete this line
	real_addr := common.BytesToAddress([]byte(in.Addr))
	is_suicide := l.svcCtx.Config.StateDB.Suicide(real_addr)

	return &sdb.SuicideResponse{IsSuicide: is_suicide}, nil
}
