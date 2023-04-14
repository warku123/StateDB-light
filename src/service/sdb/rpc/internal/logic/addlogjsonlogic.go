package logic

import (
	"context"
	"encoding/json"

	"statedbl/core/types"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogJsonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogJsonLogic {
	return &AddLogJsonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogJsonLogic) AddLogJson(in *sdb.AddLogJsonRequest) (*sdb.AddLogJsonRespond, error) {
	// todo: add your logic here and delete this line
	Log := &types.Log{}
	_ = json.Unmarshal([]byte(in.Json), Log)

	l.svcCtx.Statedb.AddLog(Log)

	return &sdb.AddLogJsonRespond{}, nil
}
