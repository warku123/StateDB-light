package logic

import (
	"context"

	"statedbl/common"
	"statedbl/core/types"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogLogic {
	return &AddLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogLogic) AddLog(in *sdb.AddLogRequest) (*sdb.AddLogRespond, error) {
	// todo: add your logic here and delete this line
	hash_list := []common.Hash{}
	for _, str := range in.Topic {
		hash_list = append(hash_list,
			common.BytesToHash([]byte(str)))
	}

	Log := &types.Log{
		Address:     common.HexToAddress(in.Addr),
		Topics:      hash_list,
		Data:        []byte(in.Data),
		BlockNumber: in.BlockNumber,
		TxHash:      common.BytesToHash([]byte(in.TxHash)),
		TxIndex:     uint(in.TxIndex),
		BlockHash:   common.BytesToHash([]byte(in.BlockHash)),
		Index:       uint(in.Index),
		Removed:     in.Removed,
	}

	l.svcCtx.Statedb.AddLog(Log)

	return &sdb.AddLogRespond{}, nil
}
