package returns

import (
	"encoding/json"
	"errors"
	"liquidation-trajectories/oracle"
	"liquidation-trajectories/pool"
	"liquidation-trajectories/tps"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

func getHfDropBounds(t []tps.UserHfRecord) (tps.HfDropBlock, error) {
	index, err := tps.ArgMax(t[:len(t)-1])
	if err != nil {
		return tps.HfDropBlock{}, err
	}
	return tps.HfDropBlock{
		User:             t[0].User,
		LiquidationBlock: t[0].LiquidationBlock,
		LowerBound:       t[index].BlockNumber,
		UpperBound:       t[index+1].BlockNumber,
	}, nil
}

func getHfDropBlock(pool *pool.Pool, db *tps.HfDropBlock) error {
	block := new(big.Int)
	block.Sub(db.UpperBound, big.NewInt(1))
	for ; block.Cmp(db.LowerBound) >= 0; block.Sub(block, big.NewInt(1)) {
		resp, err := pool.GetUserAccountData(&bind.CallOpts{BlockNumber: block}, common.HexToAddress(db.User))
		if err != nil {
			return err
		}
		if resp.HealthFactor.Cmp(big.NewInt(1000000000000000000)) > 0 {
			db.DropBlock = block
			return nil
		}
	}
	return errors.New("did not find a healthy factor between bounds")
}

func GetHfDropBlock(pool *pool.Pool, t []tps.UserHfRecord) (tps.HfDropBlock, error) {
	db, err := getHfDropBounds(t)
	if err != nil {
		return tps.HfDropBlock{}, err
	}
	err = getHfDropBlock(pool, &db)
	if err != nil {
		return tps.HfDropBlock{}, err
	}
	return db, nil
}

func GetHfDropBlocks(pool *pool.Pool, trj map[tps.LiquidationRecord][]tps.UserHfRecord) ([]tps.HfDropBlock, error) {
	var wg sync.WaitGroup
	var agg tps.HfDropBlockAggregator
	guard := make(chan struct{}, 2)
	for liq, t := range trj {
		guard <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			db, err := GetHfDropBlock(pool, t)
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
		p, err := oracle.GetAssetsPrices(&bind.CallOpts{BlockNumber: b.DropBlock}, assets)
		if err != nil {
			return make([]tps.PriceReturnRecord, 0), err
		}

		nextBlock := new(big.Int)
		nextBlock.Add(b.DropBlock, big.NewInt(1))
		np, err := oracle.GetAssetsPrices(&bind.CallOpts{BlockNumber: nextBlock}, assets)
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
