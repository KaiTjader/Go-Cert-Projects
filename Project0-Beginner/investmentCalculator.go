package main

import (
    "fmt"
    "math"
)

const inflationRate = 2.5

var investmentAmount, years float64
var expectedReturnRate = 5.5

func mainTwo() {
    fmt.Print("Investment Amount: ")
    fmt.Scan(&investmentAmount)

    fmt.Print("Expected Return Rate: ")
    fmt.Scan(&expectedReturnRate)

    futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
    futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

    // Outputs information
    fmt.Printf("Future Value: %v", futureValue)
    fmt.Println("Future Value (adjusted for inflation):", futureRealValue)
}

func calculateFutureValues() (fv float64, rfv float64) {
    fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
    rfv = fv / math.Pow(1+inflationRate/100, years)
    return
}