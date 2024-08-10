package main

import (
	"errors"
	"fmt"
	"os"
)

func mainThree() {
	revenue, revenueErr := getUserInput("Revenue: ")
	expenses, expensesErr := getUserInput("Expenses: ")
	taxRate, taxRateErr := getUserInput("Tax Rate: ")
	if revenueErr != nil || expensesErr != nil || taxRateErr != nil {
		panic(revenueErr)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f", ebt)
	fmt.Printf("%.1f", profit)
	fmt.Printf("%.3f", ratio)

	storeResults(ebt, profit, ratio)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// // Outputs information
	// fmt.Printf("Future Value: %v", futureValue)
	// fmt.Println("Future Value (adjusted for inflation):", futureRealValue)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("there was no user input")
	}
	return userInput, nil
}
