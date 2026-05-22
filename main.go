package main

import (
	"SCS/FNDS/helpers"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ltcsuite/ltcd/btcec/v2"
	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcd/ltcutil"
)

var Aj map[string]map[string]string

func updator() { // 6
	file, err := os.OpenFile("scs.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("File opening compromised | 006")
		return
	}
	dat, err2 := json.MarshalIndent(Aj, "", "  ")
	if err2 != nil {
		fmt.Println("Json compromised | 006")
		return
	}
	file.Write(dat)
	file.Sync()
	file.Close()
	fmt.Println("Sucess")

}

func commander() { //4
	for {

		fmt.Printf("Enter Command >>> ")
		scanner := bufio.NewReader(os.Stdin)
		inp, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("Command compromised | 004")
			continue
		}
		inp = strings.TrimSpace(inp)
		if strings.ToLower(inp) == "help" {
			helpers.Help()
			continue
		}
		if strings.ToLower(inp) == "newaddy" {
			helpers.Border(true)
			fmt.Printf("Enter Name >>> ")
			scanner := bufio.NewReader(os.Stdin)
			name, err := scanner.ReadString('\n')
			if err != nil {
				fmt.Println("Name compromised | 004")
				continue
			}
			name = strings.TrimSpace(name)
			NewAddress(name)
			continue
		}
		if strings.ToLower(inp) == "clear" {
			helpers.Clearer(&Aj, true)
			updator()
			continue
		}
		if strings.ToLower(inp) == "list" {
			helpers.Clearer(&Aj, false)
			continue

		}
		if strings.ToLower(inp) == "balance" {
			helpers.Balance(&Aj)
			continue
		}
		if strings.ToLower(inp) == "add" {
			name, addy, wif := helpers.Add()
			push(name, addy, wif)
		}
	}
}

func mainloader() { // 2
	Aj = make(map[string]map[string]string)

	dat, err := os.ReadFile("scs.json")
	if err != nil {
		fmt.Println("Main loading compromised Do not push i repeat do not push !!! | 002")
		return
	}
	err2 := json.Unmarshal(dat, &Aj)
	if err2 != nil {
		fmt.Println("Main loading compromised do no push i repeat do not push !!! | 002")
		return
	}

}

func existence(nm string) string {

	_, exists := Aj[nm]
	if exists {
		nm = nm + "1"
		nm = existence(nm)
		return nm

	} else {
		return nm
	}

}
func checkup(name string, address string, wif string) bool {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(address) == "" || strings.TrimSpace(wif) == "" {
		return true
	} else {
		return false
	}
}

func push(name string, address string, wif string) error { //1

	name = existence(name)
	if checkup(name, address, wif) {
		fmt.Println("Empty strings passed | 001")
		return errors.New("Red")
	}

	Aj[name] = make(map[string]string)
	Aj[name]["addy"] = address
	Aj[name]["wif"] = wif

	file, err := os.OpenFile("scs.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("File opening compromised | 001")
		return errors.New("Red")
	}
	dat, err2 := json.MarshalIndent(Aj, "", "  ")
	if err2 != nil {
		fmt.Println("Json compromised | 001")
		return errors.New("Red")
	}
	file.Write(dat)
	file.Sync()
	file.Close()
	fmt.Println("Pushed | 001 ")
	return nil

}

func NewAddress(name string) {
	pvkey, err := btcec.NewPrivateKey()
	if err != nil {
		fmt.Println("Key Creation compromised | 003")
		return
	}

	wif, err2 := ltcutil.NewWIF(pvkey, &chaincfg.MainNetParams, true)
	if err2 != nil {
		fmt.Println("Wif creation compromised | 003")
		return
	}
	publichash := ltcutil.Hash160(pvkey.PubKey().SerializeCompressed())
	addy, err3 := ltcutil.NewAddressPubKeyHash(publichash, &chaincfg.MainNetParams)
	if err3 != nil {
		fmt.Println("Addy creation compromised | 003")
		return
	}
	err4 := push(name, addy.EncodeAddress(), wif.String())
	if err4 != nil {
		fmt.Println("Addy Presistance compromised | 003")
		return
	}
	fmt.Println("Address created >>> ", addy.EncodeAddress())
	fmt.Println("wif >>> ", wif.String())

}

func main() {
	mainloader()
	commander()

}
