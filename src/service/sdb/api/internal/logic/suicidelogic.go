package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SuicideLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSuicideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SuicideLogic {
	return &SuicideLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SuicideLogic) Suicide(req *types.SuicideRequest) (resp *types.SuicideResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
