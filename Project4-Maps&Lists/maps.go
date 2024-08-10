package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// websites := []string{"google", "aws"}
	websites := map[string]string{
		"google": "google.com",
		"aws":    "amazon.com",
	}
	fmt.Println(websites["google"])
	websites["3gne"] = "3Generations.ai"

	delete(websites, "google")

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 5.0
	courseRatings["rust"] = 4.8
	courseRatings["web"] = 4

	courseRatings.output()

	for key, value := range courseRatings {
		fmt.Println(key)
		fmt.Println(value)
	}
}
