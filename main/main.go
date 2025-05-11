package main

import (
	"fmt"
	"liquidation-trajectories/oracle"
	"liquidation-trajectories/pool"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Printf("Starting Job...\n")

	// Connect to eth node
	client, err := ethclient.Dial("")
	if err != nil {
		panic(err)
	}

	// Init pool and oracle smart contracts
	pooladr := common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2")
	pool_ctr, err := pool.NewPool(pooladr, client)
	if err != nil {
		panic(err)
	}
	orcladr := common.HexToAddress("0x54586bE62E3c3580375aE3723C145253060Ca0C2")
	orl, err := oracle.NewOracle(orcladr, client)
	if err != nil {
		panic(err)
	}

	// Run ETL
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
	for ; start.Before(stop); start = start.AddDate(0, 1, 0) {
		end := start.AddDate(0, 1, 0)
		err = Etl(start, end, pool_ctr, orl, client)
		if err != nil {
			panic(err)
		}
	}
}
