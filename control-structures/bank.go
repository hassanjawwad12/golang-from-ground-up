package main

import (
	"fmt"

	"github.com/hassanjawwad12/golang-from-ground-up/filemanagement"
)

const balanceFile = "balance.txt"

func main() {

	var balance, err = filemanagement.ReadFloatFromFile(balanceFile)
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("------------")
		panic(err)
	}

	fmt.Println("Welcome to the Bank of Golang")

	for {
		presentOpions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		wantBalace := choice == 1
		wantDeposit := choice == 2
		wantWithdraw := choice == 3

		if wantBalace {
			fmt.Println("You have balance in your account", balance)
		} else if wantDeposit {
			var deposit float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scanln(&deposit)
			if deposit < 0 {
				fmt.Println("Invalid amount")
				continue
			}
			balance += deposit
			fmt.Println("You have deposited", deposit, "and your new balance is", balance)
			filemanagement.WriteFloattoFile(balance, balanceFile)
		} else if wantWithdraw {
			var withdraw float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scanln(&withdraw)
			if withdraw > balance {
				fmt.Println("Insufficient balance")
				continue
			} else {
				balance -= withdraw
				fmt.Println("You have withdrawn", withdraw, "and your new balance is", balance)
				filemanagement.WriteFloattoFile(balance, balanceFile)
			}
		} else {
			fmt.Println("Thank you for using the Bank of Golang")
			break
		}
	}
}
