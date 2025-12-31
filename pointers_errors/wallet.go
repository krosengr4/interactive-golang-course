package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bitcoin int

type Wallet struct {
	Balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	if amount <= w.Balance {
		w.Balance -= amount
	} else {
		fmt.Println("ERROR! You do not have that much to withdraw!")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("WELCOME TO YOUR BITCOIN WALLET MANAGER!")
	fmt.Println(strings.Repeat("-", 40))

	w := Wallet{}

	for {
		fmt.Println("1 - Deposit Bitcoin\n2 - Withdraw Bitcoin\n3 - See Amount\n4 - Exit")
		scanner.Scan()
		userOption, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Errorf("Invalid number: %v", err)
		}

		var userAmount int

		if userOption == 1 || userOption == 2 {
			fmt.Println("Enter the amount: $")
			scanner.Scan()
			userAmount, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Errorf("Invalid number: %v\nPlease only enter whole numbers!", err)
			}
		}

		switch userOption {
		case 1:
			w.Deposit(Bitcoin(userAmount))
		case 2:
			w.Withdraw(Bitcoin(userAmount))
		case 3:
			fmt.Println("Your Bitcoin Amount: $" + strconv.Itoa(int(w.Balance)))
		case 4:
			fmt.Println("\nSee you soon!")
			return
		}
	}

}
