package logic

import (
	"context"

	"statedbl/sdbrpc/internal/svc"
	"statedbl/sdbrpc/sdbrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *sdbrpc.Request) (*sdbrpc.Response, error) {
	// todo: add your logic here and delete this line

	return &sdbrpc.Response{}, nil
}
