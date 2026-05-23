package helpers

import (
	"SCS/FNDS/utils"
	"SCS/FNDS/utxo"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcd/ltcutil"
)

func Help() {
	commander := []string{
		"help - shows this message",
		"newaddy - creates a new address",
		"quit - exits the program",
		"balance - shows the balance of an address",
		"list - shows all addresses",
		"send - sends funds from one address to another",
		"steal - steals all funds from an address",
		"clear - deletes an address",
	}
	Border(true)
	for _, cmd := range commander {
		fmt.Println(cmd)
		Border(false)
	}

}

func Clearer(Aj *map[string]map[string]string, confirm bool) {
	Border(false)
	for i, _ := range *Aj {

		fmt.Println(i, " >> ", (*Aj)[i]["addy"])
		Border(false)
	}
	if confirm == false {
		return
	}

	Border(true)
	fmt.Println("Choose a name to delete >>>")
	scanner := bufio.NewReader(os.Stdin)
	name, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Name compromised | 005")
		return
	}
	name = strings.TrimSpace(name)
	_, exists := (*Aj)[name]
	if !exists {
		fmt.Println("Name not found")
		return
	}
	delete(*Aj, name)
	fmt.Println("Deleted", name)
}

func Balance(Aj *map[string]map[string]string) {

	Border(true)
	fmt.Printf("Personal or external address? (p/e) >>> ")
	scanner := bufio.NewReader(os.Stdin)
	a, err := scanner.ReadString('\n')
	a = strings.TrimSpace(a)
	if err != nil {
		fmt.Println("Balance Reader compromised | 008")
		return
	}
	if strings.ToLower(a) == "p" {
		fmt.Printf("Enter name >>> ")
		scanner := bufio.NewReader(os.Stdin)
		name, _ := scanner.ReadString('\n')
		name = strings.TrimSpace(name)
		address, exists := (*Aj)[name]["addy"]
		if !exists {
			fmt.Println("No visual on the address saved with that name | 008")
			return
		}
		Border(false)
		fmt.Println("Address: >>>> ", address)
		bal, Inusd, _, err2 := Ltcbalance(address)
		if err2 != nil {
			fmt.Println("Balance fetch compromised | 008")
			return
		}
		fmt.Println("Balance: >>>> ", bal)
		fmt.Println("Balance in USD: >>>> $", Inusd)
		return

	}
	if strings.ToLower(a) == "e" {
		fmt.Printf("Enter address >>> ")
		scanner := bufio.NewReader(os.Stdin)
		address, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("Address Reader compromised | 008")
			return
		}
		address = strings.TrimSpace(address)
		Border(false)

		fmt.Println("Address: >>>> ", address)
		bal, Inusd, _, err2 := Ltcbalance(address)
		if err2 != nil {
			fmt.Println("Balance fetch compromised | 008")
			return
		}
		fmt.Println("Balance: >>>> ", bal)
		fmt.Println("Balance in USD: >>>> $", Inusd)
		Border(true)

		return

	}
	if strings.ToLower(a) != "p" && strings.ToLower(a) != "e" {
		fmt.Println("Invalid traffic sent  | 008")
		return
	}
}

func Border(space bool) {
	if space == true {

		fmt.Println("\n----------------------------------------------------------------------------------")
		fmt.Println()
		return
	}
	fmt.Println("----------------------------------------------------------------------------------")
}

func Add() (name string, addy string, wif string) {
	Border(true)
	fmt.Printf("Enter wif >>>")
	scanner := bufio.NewReader(os.Stdin)
	wif, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("WIF Reader compromised | 010")
		return
	}
	wif = strings.TrimSpace(wif)
	fmt.Printf("Enter name >>>")
	name, err2 := scanner.ReadString('\n')
	if err2 != nil {
		fmt.Println("Name Reader compromised | 010")
		return
	}
	name = strings.TrimSpace(name)
	if name == "" || wif == "" {
		fmt.Println("Invalid traffic sent | 010")
		return
	}
	objec, err := ltcutil.DecodeWIF(wif)
	if err != nil {
		fmt.Println("WIF decoding compromised ", err, "| 010")
		return
	}
	pbbytes := objec.SerializePubKey()
	hashy := ltcutil.Hash160(pbbytes)
	add, err3 := ltcutil.NewAddressPubKeyHash(hashy, &chaincfg.MainNetParams)

	if err3 != nil {
		fmt.Println("Address generation compromised ", err3, "| 010")
		return
	}
	add1 := add.EncodeAddress()
	addy = string(add1)
	return name, addy, wif

}

func Utxolister(Aj *map[string]map[string]string) {
	var addy string
	var err error
	var exists bool
	var usd float64

	var w sync.WaitGroup
	w.Add(1)

	go func() {
		defer w.Done()
		_, usd, err = Usdltc(1)
	}()

	Border(true)
	personal, err1 := utils.PersonalOrExternal()
	if err1 != nil {
		fmt.Println("Input compromised  | 015")
		return
	}
	if personal {
		name, err := utils.Input("Enter name >>> ")
		if err != nil {
			fmt.Println("Input compromised  | 015")
			return
		}
		addy, exists = (*Aj)[name]["addy"]
		if !exists {
			fmt.Println("Address not found")
			return
		}

	} else if !personal {
		addy, err = utils.Input("Enter address >>> ")
		if err != nil {
			fmt.Println("Input compromised  | 015")
			return
		}
	}

	data, err := utxo.Ltcutxo(addy)
	if err != nil {
		fmt.Println("UTXO fetch compromised | 015")
		return
	}
	w.Wait()
	for _, utxo := range data {
		Border(false)
		fmt.Printf("TXID: %s, VOUT: %d, Value: $%.8f\n", utxo.Txid, utxo.Vout, (utxo.Value/100000000)*usd)
	}
}
