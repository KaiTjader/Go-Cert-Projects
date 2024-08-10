package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep((3 * time.Second))
	fmt.Println("Hello!", phrase)
}

func main() {
	// go greet("Nice to meet you")
	// go greet("How are you?")
	done := make(chan bool)
	go slowGreet(("How... are... you..."))
	// go greet("post")
}
