package main

import (
	"fmt"
	"liquidation-trajectories/apirequest"
	"liquidation-trajectories/hftrajectory"
	"liquidation-trajectories/oracle"
	"liquidation-trajectories/pool"
	"liquidation-trajectories/returns"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Printf("Starting ETL...")

	dt1 := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	dt2 := time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)

	// Extract liquidation events
	liq, err := apirequest.GetAllLiquidations(dt1, dt2, true, true)
	if err != nil {
		panic(err)
	}

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

	// Collecting users hf trajectories
	fmt.Println("STEP 1: Collecting users trajectories...")

	trajectories := hftrajectory.GetUserHfTrajectories(pool_ctr, liq)

	err = hftrajectory.SaveUserHfTrajectories(trajectories, "./outputs")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Done!\n")

	// Finding exact block where hf drops below 1 for each user
	fmt.Println("STEP 2: Collecting drop block for each liquidation event...")
	drpblocks, err := returns.GetHfDropBlocks(pool_ctr, trajectories)
	if err != nil {
		panic(err)
	}
	err = returns.SaveHfDropBlocks(drpblocks, "./outputs")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Done!\n")

	fmt.Println("STEP 3: Collecting price returns...")
	ass := []common.Address{common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")}
	prc_returns, err := returns.GetPricesReturns(orl, drpblocks, ass)
	if err != nil {
		panic(err)
	}
	err = returns.SavePriceReturns(prc_returns, "./outputs")
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}
