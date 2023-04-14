package logic

import (
	"context"

	"statedbl/common"
	"statedbl/core/types"
	"statedbl/params"
	"statedbl/service/sdb/rpc/internal/svc"
	"statedbl/service/sdb/rpc/types/sdb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrepareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrepareLogic {
	return &PrepareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrepareLogic) Prepare(in *sdb.PrepareRequest) (*sdb.PrepareRespond, error) {
	// todo: add your logic here and delete this line
	real_chain_id, err := common.StringToBig(in.Rule.ChainID)
	if err != nil {
		return &sdb.PrepareRespond{}, err
	}
	real_rules := params.Rules{
		ChainID:          real_chain_id,
		IsHomestead:      in.Rule.IsHomestead,
		IsEIP150:         in.Rule.IsEIP150,
		IsEIP155:         in.Rule.IsEIP155,
		IsEIP158:         in.Rule.IsEIP158,
		IsByzantium:      in.Rule.IsByzantium,
		IsConstantinople: in.Rule.IsConstantinople,
		IsPetersburg:     in.Rule.IsPetersburg,
		IsIstanbul:       in.Rule.IsIstanbul,
		IsBerlin:         in.Rule.IsBerlin,
		IsLondon:         in.Rule.IsLondon,
		IsMerge:          in.Rule.IsMerge,
		IsShanghai:       in.Rule.IsShanghai,
		IsCancun:         in.Rule.IsCancun,
		IsPrague:         in.Rule.IsPrague,
	}

	real_pre_compiles := []common.Address{}
	for _, str := range in.PreCompiles {
		real_pre_compiles = append(real_pre_compiles, common.HexToAddress(str))
	}

	var real_list types.AccessList
	for _, acc_tuple := range in.List {
		real_addr := common.HexToAddress(acc_tuple.Addr)
		real_storage_keys := []common.Hash{}
		for _, str_hash := range acc_tuple.StorageKeys {
			real_storage_keys = append(real_storage_keys, common.HexToHash(str_hash))
		}

		real_acc_tuple := types.AccessTuple{
			Address:     real_addr,
			StorageKeys: real_storage_keys,
		}
		real_list = append(real_list, real_acc_tuple)
	}

	real_destaddr := common.HexToAddress(in.DestAddr)

	l.svcCtx.Statedb.Prepare(
		real_rules,
		common.HexToAddress(in.SenderAddr),
		common.HexToAddress(in.CoinbaseAddr),
		&real_destaddr,
		real_pre_compiles,
		real_list,
	)

	return &sdb.PrepareRespond{}, nil
}
