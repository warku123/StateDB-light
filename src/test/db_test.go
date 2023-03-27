package db_test

import (
	"math/big"
	"path/filepath"
	"testing"

	"statedbl/common"
	"statedbl/core/rawdb"
	"statedbl/core/state"
)

func TestMemoryDB(t *testing.T) {
	// 在这先存在内存中，替代leveldb，下面测试中会测试leveldb
	mdb := rawdb.NewMemoryDatabase()
	sdb := state.NewDatabase(mdb)
	statedb, err := state.New(common.Hash{}, sdb, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	deployer := common.BytesToAddress([]byte("deployer"))

	statedb.CreateAccount(deployer)
	statedb.SetBalance(deployer, big.NewInt(1<<40))

	t.Logf("%ld", statedb.GetBalance(deployer))
	statedb.Commit(true)
}

func TestLeveldb(t *testing.T) {
	// leveldb的相关测试
	datadir := t.TempDir()
	ldb_path := filepath.Join(datadir, "geth", "chaindata")

	ldb, err := rawdb.NewLevelDBDatabase(ldb_path, 0, 0, "", false)
	if err != nil {
		t.Fatalf("%v", err)
	}
	sdb := state.NewDatabase(ldb)

	statedb, err := state.New(common.Hash{}, sdb, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	deployer := common.BytesToAddress([]byte("deployer"))

	statedb.CreateAccount(deployer)
	statedb.SetBalance(deployer, big.NewInt(1<<40))

	t.Logf("%ld", statedb.GetBalance(deployer))
	statedb.Commit(true)
	is_suicide := statedb.Suicide(deployer)
	t.Logf("%t", is_suicide)
}
