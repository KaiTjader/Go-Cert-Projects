package anonymous

import "fmt"

type transformFn func(int) int

func main() {
    numbers := []int{1, 2, 3, 4}

    double := createChanger(2)
    triple := createChanger(3)

    changed := changeNumbers(&numbers, func(number int) int {
        return number * 2
    })

    doubled := changeNumbers(&numbers, double)
    tripled := changeNumbers(&numbers, triple)

    fmt.Println(doubled)
    fmt.Println(tripled)
    fmt.Println(changed)
}

func changeNumbers(numbers *[]int, change transformFn) []int {
    dNumbers := []int{}
    for _, val := range *numbers {
        dNumbers = append(dNumbers, change(val))
    }
    return dNumbers
}

func createChanger(factor int) func(int) int {
    return func(num int) int {
        return num * factor
    }
}
