package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep((3 * time.Second))
	fmt.Println("Hellooooo!", phrase)
	doneChan <- true
	// close(doneChan)
}

func main() {
	// allDone := make([]chan bool, 4)
	done := make(chan bool)
	// done1 := allDone[0]
	// done2 := allDone[0]
	// done3 := allDone[0]
	// done4 := allDone[0]
	go greet("Nice to meet you", done)
	go greet("How are you?", done)
	go slowGreet(("How... are... you..."), done)
	go greet("post", done)
	for range done {

	}
	// for _, done := range allDone {
	// 	<-done
	// }
}
