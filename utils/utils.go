package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(prompt string) (string, error) {
	scn := bufio.NewReader(os.Stdin)
	fmt.Printf(prompt)
	inp, err := scn.ReadString('\n')
	if err != nil {
		return "", err
	}
	inp = strings.TrimSpace(inp)
	return inp, nil
}

func PersonalOrExternal() (bool, error) {
	input, err := Input("External or Personal (e/p) >>>")
	if err != nil {
		return false, err
	}
	if strings.ToLower(input) == "p" {
		return true, nil
	} else if strings.ToLower(input) == "e" {
		return false, nil
	} else {
		return false, fmt.Errorf("invalid input")
	}
}
