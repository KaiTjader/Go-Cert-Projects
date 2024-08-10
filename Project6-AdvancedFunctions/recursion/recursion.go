package recursion

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	}
	return num
	// result := 1

	// for i := 1; i <= num; i++ {
	//  result *= i
	// }

	// return result
}
