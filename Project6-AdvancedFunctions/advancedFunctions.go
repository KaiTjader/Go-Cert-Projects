package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 20, -40, 30)
	anotherSum := sumup(1, numbers...)

	fmt.Println(anotherSum)

	fmt.Println(sum)
}

func sumup(startingValue int, nums ...int) int {
	sum := startingValue

	for _, val := range nums {
		sum += val
	}
	return sum
}
