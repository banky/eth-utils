package main

import (
	"context"
	"fmt"

	"github.com/banky/eth-utils/anvil"
)

func main() {
	config := anvil.NewConfig().
		SetForkURL("https://ethereum-rpc.publicnode.com").
		SetShowLogs(true).
		SetForkBlockNumber(20000000)
	a, err := anvil.NewWithConfig(config)
	if err != nil {
		panic(err)
	}
	defer a.Close()

	// Get *ethclient.Client
	ethClient := a.EthClient()
	blockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("BlockNumber:", blockNumber)
	// Prints BlockNumber: 20000000
}
