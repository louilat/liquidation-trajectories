package apirequest

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"liquidation-trajectories/tps"
	"net/http"
	"time"
)

// Extracts the liquidation events for a given day
func GetLiquidations(date time.Time, unique bool) ([]tps.LiquidationRecord, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	dtstr := date.Format("2006-Jan-02")

	req, err := http.NewRequest("GET", "https://aavedata.lab.groupe-genes.fr/events/liquidation?date="+dtstr, nil)
	if err != nil {
		return make([]tps.LiquidationRecord, 0), err
	}

	resp, err := client.Do(req)
	if err != nil {
		return make([]tps.LiquidationRecord, 0), err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]tps.LiquidationRecord, 0), err
	}
	// fmt.Println(string(body))

	var liquidations []tps.LiquidationRecord
	e := json.Unmarshal(body, &liquidations)
	if e != nil {
		return make([]tps.LiquidationRecord, 0), e
	}

	if unique {
		liquidations = tps.GetFirstLiqPerUser(liquidations)
	}

	for i := range liquidations {
		liquidations[i].Day = date
	}
	return liquidations, nil
}

// Extracts the liquidations events over an interval of days
func GetAllLiquidations(start, stop time.Time, unique, verbose bool) ([]tps.LiquidationRecord, error) {
	var liq []tps.LiquidationRecord
	if verbose {
		fmt.Printf("Extracting liquidations from %v to %v", start, stop)
	}
	for day := start; day.Compare(stop) == -1; day = day.AddDate(0, 0, 1) {
		if verbose {
			fmt.Printf(".")
		}
		liqday, err := GetLiquidations(day, unique)
		if err != nil {
			return make([]tps.LiquidationRecord, 0), err
		}
		liq = append(liq, liqday...)
	}
	fmt.Printf("Done!\n")
	fmt.Printf("Found %v liquidation events\n", len(liq))
	return liq, nil
}
