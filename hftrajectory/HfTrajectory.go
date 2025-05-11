package hftrajectory

import (
	"encoding/json"
	"errors"
	"fmt"
	"liquidation-trajectories/pool"
	"liquidation-trajectories/tps"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

// Extracts the HF trajectory of a user over the last 24 hours before drop block.
func GetUserHfTrajectory(pool *pool.Pool, db tps.HfDropBlock) ([144]tps.UserHfRecord, error) {
	var usrrecords [144]tps.UserHfRecord
	if db.DropBlock == nil {
		return usrrecords, errors.New("drop block is nil cannot compute hf trajectory")
	}

	block := big.NewInt(0)
	delta := big.NewInt(0)
	for i := int64(0); i < 144; i++ {
		delta.Mul(big.NewInt(i), big.NewInt(50))
		block.Sub(db.DropBlock, delta)
		// fmt.Printf("Block value is %v\n", block)

		account, err := pool.GetUserAccountData(&bind.CallOpts{BlockNumber: block}, common.HexToAddress(db.User))
		if err != nil {
			return usrrecords, err
		}
		usrrecords[144-i-1] = tps.UserHfRecord{
			User:                        db.User,
			LiquidationBlock:            db.LiquidationBlock,
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

func GetUserHfTrajectories(pool *pool.Pool, db []tps.HfDropBlock) map[tps.HfDropBlock][]tps.UserHfRecord {
	var trajectories tps.UserHfRecordAggregator
	trajectories.Records = make(map[tps.HfDropBlock][]tps.UserHfRecord)

	var wg sync.WaitGroup
	guard := make(chan struct{}, 4)
	for i := range db {
		if db[i].DropBlock == nil {
			continue
		}
		wg.Add(1)
		go func() {
			guard <- struct{}{}
			defer wg.Done()
			// fmt.Println(liq[i])
			ut, err := GetUserHfTrajectory(pool, db[i])
			if err != nil {
				panic(err)
			}
			trajectories.Extend(db[i], ut[:])
			<-guard
		}()
	}
	wg.Wait()

	return trajectories.Records
}

func SaveUserHfTrajectories(trjy map[tps.HfDropBlock][]tps.UserHfRecord, path string) error {
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
