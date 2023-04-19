package svc

import (
	"log"
	"statedbl/common"
	"statedbl/core/rawdb"
	"statedbl/core/state"
	"statedbl/ethdb"
	"statedbl/service/sdb/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Statedb *state.StateDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	var bottom_db ethdb.Database

	if c.BottomDB.Is_memory {
		bottom_db = rawdb.NewMemoryDatabase()
	} else {
		bottom_db, _ = rawdb.NewLevelDBDatabase(c.BottomDB.LevelDB_dir, 0, 0, "", false)
	}

	sdb := state.NewDatabase(bottom_db)

	statedb, err := state.New(common.HexToHash(c.UpperDB.Hash), sdb, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceContext{
		Config:  c,
		Statedb: statedb,
	}
}
