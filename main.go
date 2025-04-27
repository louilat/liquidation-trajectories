package main

import (
	"liquidation-trajectories/pool"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("")
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2")
	pool_ctr, err := pool.NewPool(address, client)
	if err != nil {
		panic(err)
	}
}
