package returns

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"liquidation-trajectories/oracle"
	"liquidation-trajectories/pool"
	"liquidation-trajectories/tps"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getHfDropBounds(pool *pool.Pool, liq tps.LiquidationRecord) (tps.HfDropBlock, error) {
	block := liq.BlockNumber
	findlim := new(big.Int)
	findlim.Sub(block, big.NewInt(14400))
	for ; block.Cmp(findlim) > 0; block.Sub(block, big.NewInt(100)) {
		resp, err := pool.GetUserAccountData(&bind.CallOpts{BlockNumber: block}, common.HexToAddress(liq.User))
		if err != nil {
			return tps.HfDropBlock{}, err
		}
		if resp.HealthFactor.Cmp(big.NewInt(1000000000000000000)) > 0 {
			return tps.HfDropBlock{
				User:             liq.User,
				LiquidationBlock: liq.BlockNumber,
				LowerBound:       new(big.Int).Set(block),
				UpperBound:       new(big.Int).Set(block.Add(block, big.NewInt(100))),
			}, nil
		}
	}
	return tps.HfDropBlock{}, errors.New("could not find a healthy hf 48 hours before liquidation")
}

func getHfDropBlock(pool *pool.Pool, client *ethclient.Client, db *tps.HfDropBlock) error {
	block := new(big.Int)
	block.Sub(db.UpperBound, big.NewInt(1))
	for ; block.Cmp(db.LowerBound) >= 0; block.Sub(block, big.NewInt(1)) {
		resp, err := pool.GetUserAccountData(&bind.CallOpts{BlockNumber: block}, common.HexToAddress(db.User))
		if err != nil {
			return err
		}
		if resp.HealthFactor.Cmp(big.NewInt(1000000000000000000)) > 0 {
			b := new(big.Int)
			b.Add(block, big.NewInt(1))
			db.DropBlock = b
			header, err := client.HeaderByNumber(context.Background(), b)
			if err != nil {
				return err
			}
			db.DropBlockTime = time.Unix(int64(header.Time), 0)
			return nil
		}
	}
	return errors.New("did not find a healthy factor between bounds")
}

func GetHfDropBlock(pool *pool.Pool, client *ethclient.Client, liq tps.LiquidationRecord) (tps.HfDropBlock, error) {
	db, err := getHfDropBounds(pool, liq)
	if err != nil {
		return tps.HfDropBlock{}, err
	}
	err = getHfDropBlock(pool, client, &db)
	if err != nil {
		return tps.HfDropBlock{}, err
	}
	return db, nil
}

func GetHfDropBlocks(pool *pool.Pool, client *ethclient.Client, liq []tps.LiquidationRecord) ([]tps.HfDropBlock, error) {
	var wg sync.WaitGroup
	var agg tps.HfDropBlockAggregator
	guard := make(chan struct{}, 2)
	for _, liq := range liq {
		guard <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			db, err := GetHfDropBlock(pool, client, liq)
			if err != nil {
				agg.Extend(tps.HfDropBlock{
					User:             liq.User,
					LiquidationBlock: liq.BlockNumber,
					Error:            err.Error()})
			} else {
				agg.Extend(db)
			}
			<-guard
		}()
	}
	wg.Wait()
	return agg.Records, nil
}

func SaveHfDropBlocks(db []tps.HfDropBlock, path string) error {
	fbytes, err := json.Marshal(db)
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
	f, err := os.Create(path + "/drop_blocks.json")
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(fbytes)
	return nil
}

func GetPricesReturns(oracle *oracle.Oracle, db []tps.HfDropBlock, assets []common.Address) ([]tps.PriceReturnRecord, error) {
	var prices []tps.PriceReturnRecord
	for _, b := range db {
		// If drop block info is not provided, continue
		if b.DropBlock == nil {
			continue
		}
		block := new(big.Int)
		block.Sub(b.DropBlock, big.NewInt(300))
		p, err := oracle.GetAssetsPrices(&bind.CallOpts{BlockNumber: block}, assets)
		if err != nil {
			return make([]tps.PriceReturnRecord, 0), err
		}

		// nextBlock := new(big.Int)
		// nextBlock.Add(b.DropBlock, big.NewInt(1))
		np, err := oracle.GetAssetsPrices(&bind.CallOpts{BlockNumber: b.DropBlock}, assets)
		if err != nil {
			return make([]tps.PriceReturnRecord, 0), err
		}
		for i, r := range p {
			prices = append(prices, tps.PriceReturnRecord{
				BlockNumber:     b.DropBlock,
				UnderlyingToken: assets[i].String(),
				Price:           r,
				NextBlockPrice:  np[i],
			})
		}
	}
	return prices, nil
}

func SavePriceReturns(pr []tps.PriceReturnRecord, path string) error {
	fbytes, err := json.Marshal(pr)
	if err != nil {
		return err
	}
	f, err := os.Create(path + "/price_returns.json")
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(fbytes)
	return nil
}
