package senders

import (
	"SCS/FNDS/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcd/ltcutil"
)

func Ltcsend(Aj *map[string]map[string]string) {
	helpers.Border(true)
	fmt.Printf("External or Personal (e/p) >>>")
	s := bufio.NewReader(os.Stdin)
	inp, err := s.ReadString('\n')
	if err != nil {
		fmt.Println("Input compromised | 011")
		return
	}
	inp = strings.TrimSpace(inp)
	if strings.ToLower(inp) == "e" {
		fmt.Printf("Enter wif (sending from) >>> ")
		a := bufio.NewReader(os.Stdin)
		wif, err2 := a.ReadString('\n')
		if err2 != nil {
			fmt.Println("Input compromised | 011")
			return
		}
		wif = strings.TrimSpace(wif)
		ghostly(wif)
		return
	}

	if strings.ToLower(inp) == "p" {
		fmt.Printf("Enter name (sending from) >>> ")
		a := bufio.NewReader(os.Stdin)
		name, err2 := a.ReadString('\n')
		if err2 != nil {
			fmt.Println("Input compromised | 011")
			return
		}
		name = strings.TrimSpace(name)
		wif, exists := (*Aj)[name]["wif"]
		if !exists {
			fmt.Println("No visual on the wif saved with that name | 011")
			return
		}
		ghostly(wif)
		return

	}
	if strings.ToLower(inp) != "e" && strings.ToLower(inp) != "p" {
		fmt.Println("Invalid input | 011")
		return
	}

}

func ghostly(wif string) {
	obj, err := ltcutil.DecodeWIF(wif)
	if err != nil {
		fmt.Println("Wif decoding compromised | 013")
		return
	}
	pub := obj.PrivKey.Serialize()
	hasy := ltcutil.Hash160(pub)
	haddy, err := ltcutil.NewAddressPubKeyHash(hasy, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Address generation compromised | 013")
		return
	}

	helpers.Border(true)
	fmt.Printf("Enter destination address >>> ")
	scanner := bufio.NewReader(os.Stdin)
	dest, err2 := scanner.ReadString('\n')
	if err2 != nil {
		fmt.Println("Input compromised | 012")
		return
	}
	dest = strings.TrimSpace(dest)

	fmt.Printf("Amount in ltc or Usd (ltc/usd) >>>>> ")
	scn := bufio.NewReader(os.Stdin)
	opinion, err := scn.ReadString('\n')
	if err != nil {
		fmt.Println("Input compromised | 012")
		return
	}
	opinion = strings.TrimSpace(opinion)
	if opinion == "ltc" {
		fmt.Printf("Enter amount in ltc >>>>> ")
		lamt := bufio.NewReader(os.Stdin)
		amount, err2 := lamt.ReadString('\n')
		if err2 != nil {
			fmt.Println("Input compromised | 012")
			return
		}
		amount = strings.TrimSpace(amount)
		// to be converted into float and convert in litoshi

	}
	if strings.ToLower(opinion) == "usd" {
		fmt.Printf("Enter amount in usd >>>>> $")
		uamt := bufio.NewReader(os.Stdin)
		amount, err2 := uamt.ReadString('\n')
		if err2 != nil {
			fmt.Println("Input compromised | 012")
			return
		}
		amount = strings.TrimSpace(amount)
		// to be converted into float then converted into equivalent ltc and then converted into litoshi

	}
	if strings.ToLower(opinion) != "ltc" && strings.ToLower(opinion) != "usd" {
		fmt.Println("Invalid input | 012")
		return
	}

}
