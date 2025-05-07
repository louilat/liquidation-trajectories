package hftrajectory

import (
	"encoding/json"
	"fmt"
	"liquidation-trajectories/pool"
	"liquidation-trajectories/tps"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

// Extracts the HF trajectory of a user over the last 24 hours before liquidation.
func GetUserHfTrajectory(pool *pool.Pool, liq tps.LiquidationRecord) ([144]tps.UserHfRecord, error) {
	var usrrecords [144]tps.UserHfRecord
	block := big.NewInt(0)
	delta := big.NewInt(0)
	for i := int64(0); i < 144; i++ {
		delta.Mul(big.NewInt(i), big.NewInt(50))
		block.Sub(liq.BlockNumber, delta)
		// fmt.Printf("Block value is %v\n", block)

		account, err := pool.GetUserAccountData(&bind.CallOpts{BlockNumber: block}, common.HexToAddress(liq.User))
		if err != nil {
			return usrrecords, err
		}
		usrrecords[144-i-1] = tps.UserHfRecord{
			User:                        liq.User,
			LiquidationBlock:            liq.BlockNumber,
			BlockNumber:                 new(big.Int).Set(block),
			TotalCollateralBase:         account.TotalCollateralBase,
			TotalDebtBase:               account.TotalDebtBase,
			AvailableBorrowsBase:        account.AvailableBorrowsBase,
			CurrentLiquidationThreshold: account.CurrentLiquidationThreshold,
			Ltv:                         account.Ltv,
			HealthFactor:                account.HealthFactor,
		}
	}
	return usrrecords, nil
}

func GetUserHfTrajectories(pool *pool.Pool, liq []tps.LiquidationRecord) map[tps.LiquidationRecord][]tps.UserHfRecord {
	var trajectories tps.UserHfRecordAggregator
	trajectories.Records = make(map[tps.LiquidationRecord][]tps.UserHfRecord)

	var wg sync.WaitGroup
	guard := make(chan struct{}, 4)
	for i := range liq {
		wg.Add(1)
		go func() {
			guard <- struct{}{}
			defer wg.Done()
			// fmt.Println(liq[i])
			ut, err := GetUserHfTrajectory(pool, liq[i])
			if err != nil {
				panic(err)
			}
			trajectories.Extend(liq[i], ut[:])
			<-guard
		}()
	}
	wg.Wait()

	return trajectories.Records
}

func SaveUserHfTrajectories(trjy map[tps.LiquidationRecord][]tps.UserHfRecord, path string) error {
	var t []tps.UserHfRecord
	for _, rec := range trjy {
		t = append(t, rec...)
	}
	fbytes, err := json.Marshal(t)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
		fmt.Println("Directory created successfully.")
	} else {
		fmt.Println("Directory already exists.")
	}
	f, err := os.Create(path + "/user_trajectories.json")
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(fbytes)
	return nil
}
