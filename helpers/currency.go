package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Usdltc(ltc float64) (float64, error) { // cache system to be implemented in next commit
	intel, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=litecoin&vs_currencies=usd")
	if err != nil {
		fmt.Println("GET compromised | 009")
		return 0, err
	}
	defer intel.Body.Close()
	var data struct {
		Ltc struct {
			Usd float64 `json:"usd"`
		} `json:"litecoin"`
	}
	err2 := json.NewDecoder(intel.Body).Decode(&data)
	if err2 != nil {
		fmt.Println("Decoding compromised | 009")
		return 0, err2
	}
	return ltc * data.Ltc.Usd, nil
}
