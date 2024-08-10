package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Soccer", "Football", "Swim"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[2])
	fmt.Println(hobbies[1:])

	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	fmt.Println(len(mainHobbies))
	fmt.Println(cap(mainHobbies))

	courseGoals := []string{"learn go", "use go"}
	courseGoals[1] = "Learn all basics"
	courseGoals = append(courseGoals, "use go")
	fmt.Println(courseGoals)

	products := []Product{
		{"first-product",
			"a first prod",
			11.1},
		{"second-product",
			"jdjdj",
			11.5},
	}
	fmt.Println(products)

	newProduct := Product{
		"third-prod",
		"a third prod",
		15.99,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}
