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
)

func Etl(start, stop time.Time, pool *pool.Pool, oracle *oracle.Oracle) error {
	// Extract liquidation events
	liq, err := apirequest.GetAllLiquidations(start, stop, true, true)
	if err != nil {
		return err
	}
	// Collecting users hf trajectories
	fmt.Println("STEP 1: Collecting users trajectories...")

	trajectories := hftrajectory.GetUserHfTrajectories(pool, liq)

	err = hftrajectory.SaveUserHfTrajectories(trajectories, "./outputs")
	if err != nil {
		return err
	}
	fmt.Printf("Done!\n")

	// Finding exact block where hf drops below 1 for each user
	fmt.Println("STEP 2: Collecting drop block for each liquidation event...")
	drpblocks, err := returns.GetHfDropBlocks(pool, trajectories)
	if err != nil {
		return err
	}
	err = returns.SaveHfDropBlocks(drpblocks, "./outputs")
	if err != nil {
		return err
	}
	fmt.Printf("Done!\n")

	fmt.Println("STEP 3: Collecting price returns...")
	ass := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
	}
	prc_returns, err := returns.GetPricesReturns(oracle, drpblocks, ass)
	if err != nil {
		return err
	}
	err = returns.SavePriceReturns(prc_returns, "./outputs")
	if err != nil {
		return err
	}
	fmt.Println("Done!")
	return nil
}
