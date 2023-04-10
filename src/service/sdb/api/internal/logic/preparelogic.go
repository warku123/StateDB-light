package logic

import (
	"context"

	"statedbl/service/sdb/api/internal/svc"
	"statedbl/service/sdb/api/internal/types"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrepareLogic {
	return &PrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrepareLogic) Prepare(req *types.PrepareRequest) (resp *types.PrepareResponse, err error) {
	// todo: add your logic here and delete this line
	real_rule := &sdb.Rules{
		ChainID:          req.Rule.ChainID,
		IsHomestead:      req.Rule.IsHomestead,
		IsEIP150:         req.Rule.IsEIP150,
		IsEIP155:         req.Rule.IsEIP155,
		IsEIP158:         req.Rule.IsEIP158,
		IsByzantium:      req.Rule.IsByzantium,
		IsConstantinople: req.Rule.IsConstantinople,
		IsPetersburg:     req.Rule.IsPetersburg,
		IsIstanbul:       req.Rule.IsIstanbul,
		IsBerlin:         req.Rule.IsBerlin,
		IsLondon:         req.Rule.IsLondon,
		IsMerge:          req.Rule.IsMerge,
		IsShanghai:       req.Rule.IsShanghai,
		IsCancun:         req.Rule.IsCancun,
		IsPrague:         req.Rule.IsPrague,
	}

	real_list := []*sdb.AccessTuple{}
	for _, ele := range req.List {
		real_addr := ele.Addr
		real_skeys := ele.StorageKeys
		real_list = append(real_list, &sdb.AccessTuple{
			Addr:        real_addr,
			StorageKeys: real_skeys,
		})
	}

	_, err = l.svcCtx.SdbRpc.Prepare(l.ctx, &sdb.PrepareRequest{
		Rule:         real_rule,
		SenderAddr:   req.SenderAddr,
		CoinbaseAddr: req.CoinbaseAddr,
		DestAddr:     req.DestAddr,
		PreCompiles:  req.PreCompiles,
		List:         real_list,
	})
	return &types.PrepareResponse{}, nil
}
