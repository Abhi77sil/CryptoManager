package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		bal, Inusd, err2 := Ltcbalance(address)
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
		bal, Inusd, err2 := Ltcbalance(address)
		if err2 != nil {
			fmt.Println("Balance fetch compromised | 008")
			return
		}
		fmt.Println("Balance: >>>> ", bal)
		fmt.Println("Balance in USD: >>>> $", Inusd)

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
