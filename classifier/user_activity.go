package classifier

import (
	"fmt"
	"liquidation-trajectories/apirequest"
	"liquidation-trajectories/pool"
	"liquidation-trajectories/tps"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

// Return slice of user addresses who borrowed / withdrew / sent tokens during that day
func GetActiveUsers(date time.Time) ([]tps.ActivityRecord, error) {
	var activeusers []tps.ActivityRecord

	withdraws, err := apirequest.GetWithdrawEvents(date)
	if err != nil {
		return make([]tps.ActivityRecord, 0), err
	}
	for _, w := range withdraws {
		if !tps.Contains(activeusers, w.User, w.BlockNumber) {
			activeusers = append(activeusers, tps.ActivityRecord{User: w.User, BlockNumber: w.BlockNumber})
		}
	}

	borrows, err := apirequest.GetBorrowEvents(date)
	if err != nil {
		return make([]tps.ActivityRecord, 0), err
	}
	for _, b := range borrows {
		if !tps.Contains(activeusers, b.OnBehalfOf, b.BlockNumber) {
			activeusers = append(activeusers, tps.ActivityRecord{User: b.OnBehalfOf, BlockNumber: b.BlockNumber})
		}
	}

	balancetransfers, err := apirequest.GetBalanceTransferEvents(date)
	if err != nil {
		return make([]tps.ActivityRecord, 0), err
	}
	for _, b := range balancetransfers {
		if !tps.Contains(activeusers, b.From, b.BlockNumber) {
			activeusers = append(activeusers, tps.ActivityRecord{User: b.From, BlockNumber: b.BlockNumber})
		}
	}
	return activeusers, nil
}

// Return true if user did some actions 24h before drop block, that directly led to hf < 1.05, else false
func WasUserActive24hBeforeDb(pool *pool.Pool, db tps.HfDropBlock, activity []tps.ActivityRecord, hftrjy []tps.UserHfRecord) (bool, error) {
	// Find user activities 24h before liq and strictly before liq
	user := hftrjy[0].User
	var useract []tps.ActivityRecord
	for _, a := range activity {
		if a.User == user && a.BlockNumber.Cmp(hftrjy[0].BlockNumber) >= 0 && a.BlockNumber.Cmp(db.DropBlock) < 0 {
			useract = append(useract, a)
		}
	}
	if len(useract) == 0 {
		// If no activity, return false
		return false, nil
	} else {
		// If user has some activities, check if block just after liq has hf <= 1.05
		for _, a := range useract {
			resp, err := pool.GetUserAccountData(&bind.CallOpts{BlockNumber: a.BlockNumber}, common.HexToAddress(user))
			if err != nil {
				return false, err
			}
			if resp.HealthFactor.Cmp(big.NewInt(1050000000000000000)) <= 0 {
				return true, nil
			}
			// nexthf, err := tps.MinAfterBlock(hftrjy, a.BlockNumber)
			// if err != nil {
			// 	return false, err
			// }
			// if nexthf.HealthFactor.Cmp(big.NewInt(1050000000000000000)) <= 0 {
			// 	return true, nil
			// }
		}
	}
	return false, nil
}

// Set UserCategory field of liquidations records to 3 for users
// who borrwed / withdrew / sent atokens 24 hours before liquidation
func FindUserCategory(pool *pool.Pool, trajectories map[tps.HfDropBlock][]tps.UserHfRecord) ([]tps.HfDropBlock, error) {
	var dropblocks []tps.HfDropBlock
	var actcache tps.ActiveUsersAggregator
	actcache.Activity = make(map[int64][]tps.ActivityRecord)

	for db, trjy := range trajectories {
		t := time.Date(db.DropBlockTime.Year(), db.DropBlockTime.Month(), db.DropBlockTime.Day(), 0, 0, 0, 0, time.UTC)
		_, exists := actcache.Activity[t.Unix()]
		if !exists {
			act, err := GetActiveUsers(t)
			if err != nil {
				return make([]tps.HfDropBlock, 0), err
			}
			actcache.Extend(t.Unix(), act)
		}
		_, exists = actcache.Activity[t.AddDate(0, 0, -1).Unix()]
		if !exists {
			prevact, err := GetActiveUsers(t.AddDate(0, 0, -1))
			if err != nil {
				return make([]tps.HfDropBlock, 0), err
			}
			actcache.Extend(t.AddDate(0, 0, -1).Unix(), prevact)
		}
		activity := actcache.Activity[t.Unix()]
		prevactivity := actcache.Activity[t.AddDate(0, 0, -1).Unix()]
		activity = append(prevactivity, activity...)

		wasactive, err := WasUserActive24hBeforeDb(pool, db, activity, trjy)
		if err != nil {
			return make([]tps.HfDropBlock, 0), err
		}

		if wasactive {
			fmt.Println("was active")
			db.UserCategory = 3
		} else if tps.MaxHealthFactor(trjy).Cmp(big.NewInt(1200000000000000000)) >= 0 {
			db.UserCategory = 1
		} else if tps.MaxHealthFactor(trjy).Cmp(big.NewInt(1100000000000000000)) >= 0 {
			db.UserCategory = 2
		} else {
			db.UserCategory = 4
		}
		dropblocks = append(dropblocks, db)
	}
	return dropblocks, nil
}
