package main

import (
	"fmt"

	"example.com/files/fileOps"
	"github.com/Pallinder/go-randomdata"
)
// This route works when its outside a file

const accountBalanceFileName = "balance.txt"

func main() {
	var accountBalance, err = fileOps.GetFloatFromFile(accountBalanceFileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		panic(err)
	}

	fmt.Println("Welcome to the Bank!")
	fmt.Println("Reach us 24-7: ", randomdata.PhoneNumber())
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Balance amount: %.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Printf("How much would you like to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid answer, must be greater than 0!")
				continue
			}
			accountBalance += depositAmount

			fmt.Printf("Completed! New balance: %.2f\n", accountBalance)
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFileName)
		case 3:
			var withdrawAmount float64
			fmt.Printf("How much would you like to deposit: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid answer, must be greater than 0!")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount

			fmt.Printf("Completed! New balance: %.2f\n", accountBalance)
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFileName)
		default:
			fmt.Println("Thanks for choosing this bank")
			return
		}
	}
}
