package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation float64 = 2.5
	var invAmount float64
	var years, err float64 = 5.5, 10

	fmt.Scan(&invAmount)

	preCalc := invAmount * math.Pow(1+err/100, years)
	output := preCalc / math.Pow(1+inflation/100, years)
	fmt.Print(output)
}
