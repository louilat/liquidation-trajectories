package main

import (
	"fmt"
	"liquidation-trajectories/apirequest"
	"liquidation-trajectories/classifier"
	"liquidation-trajectories/hftrajectory"
	"liquidation-trajectories/oracle"
	"liquidation-trajectories/pool"
	"liquidation-trajectories/returns"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Etl(start, stop time.Time, pool *pool.Pool, oracle *oracle.Oracle, client *ethclient.Client) error {
	date_str := start.String()[:7]
	fmt.Printf("Starting ETL for month %v...\n", date_str)
	// Extract liquidation events
	liq, err := apirequest.GetAllLiquidations(start, stop, true, true)
	if err != nil {
		return err
	}
	// Collecting drop block for each liquidation event
	fmt.Println("STEP 1: Collecting drop blocks...")

	dropblocks, err := returns.GetHfDropBlocks(pool, client, liq)
	if err != nil {
		return err
	}
	fmt.Printf("Done!\n")

	fmt.Println("STEP 2: Collecting price returns...")
	assets := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
	}
	prc_returns, err := returns.GetPricesReturns(oracle, dropblocks, assets)
	if err != nil {
		return err
	}
	err = returns.SavePriceReturns(prc_returns, "./outputs2/outputs_"+date_str)
	if err != nil {
		return err
	}
	fmt.Printf("Done!\n")

	// Collecting user hf trajectory
	fmt.Println("STEP 3: Collecting hf trajectory 24h before drop block...")
	hf_trjy := hftrajectory.GetUserHfTrajectories(pool, dropblocks)
	err = hftrajectory.SaveUserHfTrajectories(hf_trjy, "./outputs2/outputs_"+date_str)
	if err != nil {
		return err
	}
	fmt.Printf("Done!\n")

	dropblocks_activity, err := classifier.FindUserCategory(pool, hf_trjy)
	if err != nil {
		panic(err)
	}

	err = returns.SaveHfDropBlocks(dropblocks_activity, "./outputs2/outputs_"+date_str)
	if err != nil {
		return err
	}
	return nil
}
