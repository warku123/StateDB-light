package main

import (
	"fmt"
	"statedbl/common"
)

func main() {
	fmt.Println(common.BytesToAddress([]byte("deployer")))
}
