package main

import "fmt"

func main() {
	age := 30 // Regular variable

	agePointer := &age
	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(&age)
	fmt.Println("Adult years:", age)

	// adultYears := getAdultYears(age)
	// fmt.Println(adultYears)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
