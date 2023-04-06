package svc

import (
	"statedbl/service/sdb/api/internal/config"
	"statedbl/service/sdb/rpc/sdbclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	SdbRpc sdbclient.Sdb
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SdbRpc: sdbclient.NewSdb(zrpc.MustNewClient(c.SdbRpc)),
	}
}
