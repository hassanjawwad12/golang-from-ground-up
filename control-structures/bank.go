package main

import "fmt"

func main() {
	var balance float64 = 1000.50

	fmt.Println("Welcome to the Bank of Golang")

	for {
		fmt.Println("What do u want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

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
			}
		} else {
			fmt.Println("Thank you for using the Bank of Golang")
			break
		}
	}

}
