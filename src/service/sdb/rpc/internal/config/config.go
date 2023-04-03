package config

import (
	"statedbl/core/state"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	BottomDB struct {
		Is_memory   bool
		LevelDB_dir string
	}

	UpperDB struct {
		Hash string
	}

	StateDB *state.StateDB
}
