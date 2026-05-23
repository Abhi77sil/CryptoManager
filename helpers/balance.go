package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func Ltcbalance(addr string) (float64, float64, float64, error) {
	var w sync.WaitGroup
	var err3 error
	var Inusd float64
	w.Add(1)
	go func() {
		defer w.Done()
		Inusd, _, err3 = Usdltc(1)
	}()

	var data struct {
		Info struct {
			Recived float64 `json:"funded_txo_sum"`
			Sent    float64 `json:"spent_txo_sum"`
		} `json:"chain_stats"`
	}
	traffic, err := http.Get("https://litecoinspace.org/api/address/" + addr)
	if err != nil {
		fmt.Println("GET compromised | 007")
		return 0, 0, 0, err
	}
	defer traffic.Body.Close()
	err2 := json.NewDecoder(traffic.Body).Decode(&data)
	if err2 != nil {
		fmt.Println("Decoding compromised | 007")
		return 0, 0, 0, err2
	}
	raw := data.Info.Recived - data.Info.Sent
	balance := (data.Info.Recived - data.Info.Sent) / 100000000

	if err3 != nil {
		fmt.Println("USD conversion compromised | 007")
		return balance, 0, 0, err3
	}
	w.Wait()
	return balance, balance * Inusd, raw, nil

}
