package utxo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Ltcutxo(addy string) (returner []struct {
	Txid  string  `json:"txid"`
	Vout  int     `json:"vout"`
	Value float64 `json:"value"`
}, err error) {
	var data []struct {
		Txid  string  `json:"txid"`
		Vout  int     `json:"vout"`
		Value float64 `json:"value"`
	}
	resp, err := http.Get("https://litecoinspace.org/api/address/" + addy + "/utxo")
	if err != nil {
		fmt.Println("Get compromised | 014")
		return nil, err
	}
	defer resp.Body.Close()
	err2 := json.NewDecoder(resp.Body).Decode(&data)
	if err2 != nil {
		fmt.Println("Decode compromised | 015")
		return nil, err2
	}
	return data, nil
}

func Chooseutxo(price float64, addy string) {
	utxos, err := Ltcutxo(addy)
	if err != nil {
		fmt.Println("Utxo retrieval compromised | 016")
		return
	}

}
