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
	}
	fmt.Println("\n \n ---------------------------------------------------------------------------")
	for _, cmd := range commander {
		fmt.Println(cmd)
		fmt.Println("----------------------------------------------------------------------------")
	}

}

func Clearer(Aj *map[string]map[string]string, confirm bool) {
	fmt.Println("----------------------------------------------------------------")
	for i, _ := range *Aj {

		fmt.Println(i, " >> ", (*Aj)[i]["addy"])
		fmt.Println("----------------------------------------------------------------")
	}
	if confirm == false {
		return
	}

	fmt.Println("-----------------------------------")
	fmt.Println("Choose a name to delete >>>")
	scanner := bufio.NewReader(os.Stdin)
	name, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Name compromised | 005")
		return
	}
	name = strings.TrimSpace(name)
	for key := range *Aj {
		if key == name {
			delete(*Aj, key)
			fmt.Println("Deleted", key)
			return
		}
	}
	fmt.Println("Name not found")

}
