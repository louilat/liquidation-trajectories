package tps

import (
	"math/big"
	"slices"
	"time"
)

type LiquidationRecord struct {
	Day                        time.Time `json:"day"`
	BlockNumber                *big.Int  `json:"blockNumber"`
	CollateralAsset            string    `json:"collateralAsset"`
	DebtAsset                  string    `json:"debtAsset"`
	User                       string    `json:"user"`
	DebtToCover                *big.Int  `json:"debtToCover"`
	LiquidatedCollateralAmount *big.Int  `json:"liquidatedCollateralAmount"`
	Liquidator                 string    `json:"liquidator"`
}

// Returns true if block number of hfrec is lower than block number of r, else false
func (liqrec LiquidationRecord) IsLess(r LiquidationRecord) bool {
	return liqrec.BlockNumber.Cmp(r.BlockNumber) < 0
}

// Returns the record with the lowest block among a slice of records
func min(r []LiquidationRecord) LiquidationRecord {
	min := r[0]
	for _, rec := range r {
		if rec.IsLess(min) {
			min = rec
		}
	}
	return min
}

// Returns slice of unique users inside slice of liquidation records
func getusers(r []LiquidationRecord) []string {
	var allusr []string
	for _, rec := range r {
		if !slices.Contains(allusr, rec.User) {
			allusr = append(allusr, rec.User)
		}
	}
	return allusr
}

// Filter a liquidation records slice that correspond to specific user
func filter(r []LiquidationRecord, usr string) []LiquidationRecord {
	var fr []LiquidationRecord
	for _, rec := range r {
		if rec.User == usr {
			fr = append(fr, rec)
		}
	}
	return fr
}

// Return first liquidation for each user in r
func GetFirstLiqPerUser(r []LiquidationRecord) []LiquidationRecord {
	var uniqueliq []LiquidationRecord
	users := getusers(r)
	for _, usr := range users {
		usrliqrec := min(filter(r, usr))
		uniqueliq = append(uniqueliq, usrliqrec)
	}
	return uniqueliq
}
