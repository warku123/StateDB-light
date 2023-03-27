package db_test

import (
	"math/big"
	"testing"

	"statedbl/common"
	"statedbl/core/rawdb"
	"statedbl/core/state"
)

func TestCreateSet(t *testing.T) {
	// 在这先存在内存中，替代leveldb，下面测试中会测试leveldb
	rdb := rawdb.NewMemoryDatabase()
	sdb := state.NewDatabase(rdb)
	statedb, _ := state.New(common.Hash{}, sdb, nil)

	deployer := common.BytesToAddress([]byte("deployer"))

	statedb.CreateAccount(deployer)
	statedb.SetBalance(deployer, big.NewInt(1<<40))

	t.Logf("%ld", statedb.GetBalance(deployer))
}
