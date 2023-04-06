package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	BottomDB struct {
		Is_memory   bool
		LevelDB_dir string
	}

	UpperDB struct {
		Hash string
	}

	SdbRpc zrpc.RpcClientConf
}
