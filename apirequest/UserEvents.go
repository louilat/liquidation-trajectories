package apirequest

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"liquidation-trajectories/tps"
	"net/http"
	"time"
)

// Extracts all the withdraw events for a given day
func GetWithdrawEvents(date time.Time) ([]tps.WithdrawEvent, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	dtstr := date.Format("2006-Jan-02")

	req, err := http.NewRequest("GET", "https://aavedata.lab.groupe-genes.fr/events/withdraw?date="+dtstr, nil)
	if err != nil {
		return make([]tps.WithdrawEvent, 0), err
	}

	resp, err := client.Do(req)
	if err != nil {
		return make([]tps.WithdrawEvent, 0), err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]tps.WithdrawEvent, 0), err
	}
	// fmt.Println(string(body))

	var withdraws []tps.WithdrawEvent
	e := json.Unmarshal(body, &withdraws)
	if e != nil {
		return make([]tps.WithdrawEvent, 0), e
	}

	return withdraws, nil
}

// Extracts all the borrow events for a given day
func GetBorrowEvents(date time.Time) ([]tps.BorrowEvent, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	dtstr := date.Format("2006-Jan-02")

	req, err := http.NewRequest("GET", "https://aavedata.lab.groupe-genes.fr/events/borrow?date="+dtstr, nil)
	if err != nil {
		return make([]tps.BorrowEvent, 0), err
	}

	resp, err := client.Do(req)
	if err != nil {
		return make([]tps.BorrowEvent, 0), err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]tps.BorrowEvent, 0), err
	}
	// fmt.Println(string(body))

	var borrows []tps.BorrowEvent
	e := json.Unmarshal(body, &borrows)
	if e != nil {
		return make([]tps.BorrowEvent, 0), e
	}

	return borrows, nil
}

// Extracts all the aToken transfer events for a given day
func GetBalanceTransferEvents(date time.Time) ([]tps.TransferATokenEvent, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	dtstr := date.Format("2006-Jan-02")

	req, err := http.NewRequest("GET", "https://aavedata.lab.groupe-genes.fr/events/balancetransfer?date="+dtstr, nil)
	if err != nil {
		return make([]tps.TransferATokenEvent, 0), err
	}

	resp, err := client.Do(req)
	if err != nil {
		return make([]tps.TransferATokenEvent, 0), err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]tps.TransferATokenEvent, 0), err
	}
	// fmt.Println(string(body))

	var transfers []tps.TransferATokenEvent
	e := json.Unmarshal(body, &transfers)
	if e != nil {
		return make([]tps.TransferATokenEvent, 0), e
	}

	return transfers, nil
}
